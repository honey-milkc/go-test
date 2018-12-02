# Sync
## sync.Mutex
여러 고루틴에서 공유하는 데이터 보호
* 지원하는 Method
  ```go
  func (rw *Mutex) Lock()
  func (rw *Mutex) Unlock()
  ```

## sync.RWMutex
읽기 동작과 쓰기 동작을 나누어 잠금 처리
1. 읽기 잠금 
   * 읽기 작업에 한해서 공유 데이터가 변하지 않음을 보장. 즉, 읽기 잠금으로 잠금 처리해도 다른 고루틴에서 잠금 처리한 데이터를 읽을 수는 있지만, 변경 불가
   * 제공하는 Method
     ```go
     func (rw *RWMutex) RLock()
     func (rw *RWMutex) RUnlock()
     ```
2. 쓰기 잠금 
   * 공유 데이터 쓰기 작업을 보장하기 위한 것. 쓰기 잠금으로 잠금 처리하면 다른 고루틴에서 읽기와 쓰기 작업 모두 할 수 없다.
   * 제공하는 Method
     ```go
     func (rw *RWMutex) Lock()
     func (rw *RWMutex) Unlock()
     ```

## sync.Once
특정 함수를 한 번만 수행해야 할 때!
* 제공하는 Method
  ```go
  func (o *Once) Do(f func())
  ```
* 예제
  ```go
  c.once.Do(func() {
    c.i = initalValue
  })
  ```

## sync.WaitGroup
모든 고루틴이 종료될 때까지 대기 [[예제](https://github.com/honey-milkc/go-test/blob/master/book1/chapter5/sync.go#L100)]
* 제공하는 Method
  ```go
  func (wg *WaitGroup) Add(delta int) 
  func (wg *WaitGroup) Done()
  func (wg *WaitGroup) Wait()
  ```
* Add()
  * WaitGroup에 대기 중인 고루틴 개수 추가
* Done()
  * 대기 중인 고루틴의 수행이 종료되는 것을 알려줌
* Wait()
  * 모든 고루틴이 종료될 때까지 대기
* 주의 사항!
  * Add로 추가한 고루틴의 개수와 Done 호출한 횟수가 같아야 한다는 것!!!! 
    Wait을 호출하면 대기 중인 모든 고루틴이 종료될 때까지 대기하므로!  


