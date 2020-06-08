# Concurrency 

## Concurrency Lock

* 锁为成所需原提供了最小程度的调度控制，线程为程序员创建的实体，但是被操作系统调度，
  锁让程序员获得一些**控制权**，锁将操作系统调度的混乱状态变得更为可控
  
## Pthread锁
  
* POSIX库将锁称为互斥量**mutex**，用来提供线程之间的互斥
  
```C
pthread_mutex_t lock = PTHREAD_MUTEX_INITIALIZER;

Pthread_mutex_lock(&lock);      // wrapper for pthread_mutex_lock()
balance = balance + 1;
Pthread_mutex_unlock(&lock);
```

## 评价锁

* 提供互斥性 （mutual exclusion）
* 公平性 （fairness）
* 性能 （performance）

## Concurrency Data Structure

## Semaphore 信号量

### Definition 定义

```c
sem_t m;
sem_init(&m, 0, 1); 

sem_wait(&m);
// critical section here
sem_post(&m)
```

### Reader-Writer lock 读写锁

```c
typedef struct _rwlock_t {
    sem_t lock;
    sem_t writelock;
    int readers;
} rwlock_t


void rw_lock_init(rwlock_t *rw) {
    rw->readers = 0;
    sem_init(&rw->lock, 0, 1);
    sem_init(&rw->writelock, 0, 1);
}

void rwlock_acquire_readlock(rwlock_t *rw) {
    sem_wait(&rw->lock);
    rw->readers++;
    if (rw->readers == 1)
        sem_wait(&rw->writelock); // first reader acquires writelock
    sem_post(&rw->lock);
}

void rwlock_release_readlock() {
    sem_wait(&rw->lock);
    rw->readers--;
    if (rw->readers == 0) 
        sem_post(&rw->writelock); // last reader releases writelock
    sem_post(&rw->lock);
}

void rwlock_acquire_writelock() {
    sem_wait(&rw->writelock);
}

void rwlock_release_writelock() {
    sem_post(&rw->writelock);
}

```