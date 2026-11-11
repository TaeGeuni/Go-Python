# 3736-Minimum-Moves-to-Equal-Array-Elements-III

## 💡 문제 링크
* [LeetCode - 문제 이름](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-iii)

## 🧠 문제 접근 및 풀이
* **알고리즘 분류:** math
* **시간 복잡도:** O(N)
* **공간 복잡도:** O()

### 핵심 아이디어

특별히 새로운것을 배웠다 보다는 배열의 값의 전체합의 값을 구해
최대의 경우의 값 - 배열 전체 값을 이용해 답을 구하는것이 신기했다.

그리고 go 코드에서 반복문을 구현할때
`for _, v := range nums`
이런 형식으로 반복문을 작성하는게 go의 스타일 인것을 알았다.
`for i, v := range nums`
참고로 이렇게 작성하면 i로 몇번째 반복문인지 확인가능.