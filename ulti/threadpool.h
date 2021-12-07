#include <bits/stdc++.h>
#include <condition_variable>
#include <functional>
#include <future>
#include <iostream>
#include <memory>
#include <mutex>
#include <thread>
#include <utility>

namespace ulti {
class ThreadPool {
  enum class ThreadState {
    ThreadInit = 0,
    ThreadWaiting = 1,
    ThreadRunning = 2,
    ThreadStop = 3
  };
  enum class ThreadType { Origin = 0, Core = 1, Cache = 2 };
  using CacheThreadWait = std::chrono::milliseconds;
  using ThreadPtr = std::shared_ptr<std::thread>;
  using ThreadId = std::atomic<int>;
  using ThreadStateAtomic = std::atomic<ThreadState>;
  using ThreadTypeAtomic = std::atomic<ThreadType>;

  struct ThreadPoolConfig {
    int CoreThreads;
    int MaxThreads;
    int MaxTasks;
    CacheThreadWait Timeout;
  };

  struct ThreadWrapper {
    ThreadPtr Ptr;
    ThreadId Id;
    ThreadStateAtomic State;
    ThreadTypeAtomic Type;

    ThreadWrapper() {
      Ptr = NULL;
      Id = 0;
    }
  };
  using ThreadWrapperPtr = std::shared_ptr<ThreadWrapper>;

  ThreadPool(const ThreadPoolConfig &config) : config_(config) {
    this->total_function_num_.store(0);
    this->waiting_thread_num_.store(0);
    this->thread_id_.store(0);
    this->is_shutdown_.store(false);
    this->is_shutdown_now_.store(false);
  }

  ~ThreadPool() {}

  bool ResetConfig(const ThreadPoolConfig& config) {
    if (config.CoreThreads < 1 || config.Timeout.count() < 1 || config.MaxThreads < config.CoreThreads) 
      return false;
    return true;
  }

  template<typename F, typename... Args>
  auto AddTask(F&& f, Args&&... args) -> std::shared_ptr<std::future<std::result_of<F(Args...)>>> {
    if (this->is_shutdown_now_.load() || this->is_shutdown_.load() || !IsValidConfig())
      return nullptr;
    
  }
private:
  ThreadPoolConfig config_;
  std::vector<ThreadWrapperPtr> worker_threads_;
  std::queue<std::function<void()>> tasks_;
  std::mutex task_mutex_;
  std::condition_variable task_cv_;
  std::atomic<int> total_function_num_;
  std::atomic<int> waiting_thread_num_;
  std::atomic<int> thread_id_;
  std::atomic<bool> is_shutdown_now_;
  std::atomic<bool> is_shutdown_;
  std::atomic<bool> is_available_;

  bool IsValidConfig() {
    if (config_.CoreThreads < 1 || config_.Timeout.count() < 1 ||
        config_.MaxThreads < config_.CoreThreads)
      return false;
    return true;
  }
};
}
/*
分为Cache线程和Core线程，Cache线程可以动态加，如果timeout取不到任务，终止线程
*/
/*
线程池分为动态和静态两种，
静态线程池在初始化的时候就分配好工作线程数量，
动态线程一开始也分配一定数量的线程，如果需要动态加入工作线程，运行不同的逻辑，如果工作队列为空，那么动态线程销毁

*/