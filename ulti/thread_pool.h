#pragma once

#include <algorithm>
#include <condition_variable>
#include <mutex>
#include <thread>
#include <functional>
#include <iostream>
#include <cstdio>
#include <vector>
#include <queue>
#include <atomic>

template<typename TaskT>
class ThreadPool {
public:
    explicit ThreadPool(int cnt_threads = 0);
    ~ThreadPool();
    void AddTask(const TaskT &task);
    void AddTask(TaskT &&task);
    int GetPendingTaskCnt() const {return m_tasks.size();}
    int GetRunningTaskCnt() const {return m_running_task_cnt;}
    void WaitAllTaskDone();
private:
    void WorkerLoop();

    std::vector<std::thread> m_thread_pool;
    int m_cnt_max_threads;
    std::condition_variable m_cond, m_cond_over;
    std::queue<TaskT> m_tasks;
    std::mutex m_mutex, m_mutex_over;
    bool m_exit;
    std::atomic_int m_running_task_cnt;
};

template <typename TaskT>
void ThreadPool<TaskT>::WaitAllTaskDone() {
    std::unique_lock<std::mutex> locker(m_mutex_over);
    while (m_running_task_cnt > 0 || m_tasks.size() > 0) {
        m_cond_over.wait(locker);
    }
}

template <typename TaskT>
ThreadPool<TaskT>::ThreadPool(int cnt_threads) {
    m_exit = false;
    m_running_task_cnt = 0;
    if (cnt_threads <= 0 ) {
        m_cnt_max_threads = std::thread::hardware_concurrency();
    } else {
        m_cnt_max_threads = cnt_threads;
    }
    for (size_t i = 0; i < m_cnt_max_threads; i++) {
        m_thread_pool.push_back(std::thread(&ThreadPool<TaskT>::WorkerLoop, this));
        m_thread_pool[i].detach();
    }
}

template <typename TaskT>
ThreadPool<TaskT>::~ThreadPool() {
    m_exit = true;
    std::unique_lock<std::mutex> locker(m_mutex_over);
    while (m_running_task_cnt > 0){
        m_cond_over.wait(locker);
    }
}

template <typename TaskT>
void ThreadPool<TaskT>::WorkerLoop() {
    while (!m_exit) {
        std::unique_lock<std::mutex> locker(m_mutex);
        while (m_tasks.empty() && !m_exit) {
            m_cond.wait(locker);
        }
        if (m_tasks.empty()) {
            break;
        }
        TaskT thread_job = std::move(m_tasks.front());
        m_running_task_cnt++;
        m_tasks.pop();
        locker.unlock();
        thread_job();
        m_running_task_cnt--;
        m_cond_over.notify_all();
    }
}

template<typename TaskT>
void ThreadPool<TaskT>::AddTask(const TaskT &task) {
    {
        std::lock_guard<std::mutex> locker(m_mutex);
        m_tasks.push(task);
    }
    m_cond.notify_one();
}

template<typename TaskT>
void ThreadPool<TaskT>::AddTask(TaskT &&task) {
    {
        std::lock_guard<std::mutex> locker(m_mutex);
        m_tasks.push(std::forward<TaskT>(task));
    }
    m_cond.notify_one();
}