import requests

# GitHub API URL
url = "https://api.github.com/repos/only9464/Fuck-YuKeTang/releases/latest"

# 发送请求
response = requests.get(url)

# 解析 JSON 响应
data = response.json()

# 提取 tag 名称
latest_tag = data.get("tag_name")

print(f"Latest tag: {latest_tag}")
