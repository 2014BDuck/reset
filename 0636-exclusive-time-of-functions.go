class Solution:
    def exclusiveTime(self, n: int, logs: List[str]) -> List[int]:
        stack = []
        result = [0 for _ in range(n)]
        prev_time  = 0
        for log in logs:
            task, op, idx = log.split(":")
            task, idx = int(task), int(idx)
            if op == "start":
                if stack:
                    result[stack[-1]] += idx - prev_time
                stack.append(task)
                prev_time = idx
            else:
                result[stack.pop()] += idx - prev_time + 1
                prev_time = idx + 1
        return result
    
