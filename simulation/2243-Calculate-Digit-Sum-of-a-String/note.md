# 2243-Calculate-Digit-Sum-of-a-String

## 💡 문제 링크
* [LeetCode - 2243-Calculate-Digit-Sum-of-a-String](https://leetcode.com/problems/calculate-digit-sum-of-a-string)

## 🧠 문제 접근 및 풀이
* **알고리즘 분류:** simulation
* **시간 복잡도:** O()
* **공간 복잡도:** O()

### 핵심 아이디어
이 문제는 버퍼를 이용해 반복문으로 계속 문자열을 저장할때 사용하는 메모리를 절약하는것이 포인트.

go의 경우는 이렇게 선언하고

```go
var buffer strings.Builder
```

파이썬은 이렇게 선언
```python
buffer = []
```

