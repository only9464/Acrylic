<template>
  <div class="main-container">
    <h1>设置界面</h1>
    <div class="settings-options">

      <!-- 选项0 -->
      <div class="setting-option">
        <span class="option-label">当前版本号：</span>
        <div class="option-control">
            <span>1.0.0( 最新版本：{{ latestVersion }} )</span>
        </div>
      </div>

      <!-- 选项1 -->
      <div class="setting-option">
        <span class="option-label">开启通知</span>
        <div class="option-control">
          <!-- 未来的控件，例如开关按钮 -->
          <input type="checkbox" id="notifications" @change="toggleNotifications" />
          <label for="notifications" class="switch"></label>
        </div>
      </div>

      <!-- 选项2 -->
      <div class="setting-option">
        <span class="option-label">选择下载路径</span>
        <div class="option-control">
          <!-- 未来的控件，例如路径选择 -->
          <input type="text" v-model="downloadPath" placeholder="请输入路径" />
        </div>
      </div>

      <!-- 选项3 -->
      <div class="setting-option">
        <span class="option-label">用户描述</span>
        <div class="option-control">
          <!-- 未来的控件，例如字符串描述 -->
          <textarea v-model="userDescription" placeholder="请输入描述"></textarea>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

// 定义响应式变量
const latestVersion = ref('')

// 定义执行加载时的异步函数
const executeOnLoad = async () => {
  try {
    console.log('设置页面已加载，执行初始化函数')
    
    // 调用 Go 函数并等待其完成
    await window.go.Settings.App.Settings(1, 2)
    
    // 异步获取最新的版本号，并设置到 latestVersion
    const version = await window.go.Settings.App.Get_latest_version_code()
    latestVersion.value = version

    console.log("window.go.Settings.App", window.go.Settings.App)
    console.log("latestVersion", latestVersion.value)
    
    // 其他初始化逻辑
  } catch (error) {
    console.error('执行初始化函数时出错:', error)
  }
}

onMounted(() => {
  executeOnLoad()
})

// 其他响应式数据和方法
const downloadPath = ref('')
const userDescription = ref('')

const toggleNotifications = (event) => {
  const isEnabled = event.target.checked
  console.log(`通知已${isEnabled ? '开启' : '关闭'}`)
  // 添加处理逻辑
}

const handleDownloadPath = () => {
  console.log(`下载路径: ${downloadPath.value}`)
  // 添加处理逻辑
}

const handleUserDescription = () => {
  console.log(`用户描述: ${userDescription.value}`)
  // 添加处理逻辑
}
</script>




<style scoped>
.main-container {
  padding: 0px;
}

.settings-options {
  display: flex;
  flex-direction: column;
  gap: 16px;
  margin-top: 20px;
}

.setting-option {
  display: flex;
  justify-content: space-between;
  align-items: center;
  /* 复用毛玻璃圆角框效果 */
  color: #000000;
  padding: 10px 15px; /* 减少左右内边距以靠近左边 */
  border-radius: 8px;
  font-weight: 600;
  letter-spacing: 0.5px;
  text-shadow: 0 1px 1px rgba(255, 255, 255, 0.349);
  background: linear-gradient(
    90deg,
    rgba(173, 216, 230, 0.4),  /* 粉碎冰蓝色调 */
    rgba(255, 182, 193, 0.6)   /* 玫瑰色调 */
  );
  backdrop-filter: blur(20px) saturate(125%);
  -webkit-backdrop-filter: blur(20px) saturate(125%);
  border: 1px solid rgba(255, 255, 255, 0.9);
  box-shadow: 0 8px 32px 0 rgba(0, 0, 0, 0.15);
  cursor: default; /* 修改为默认光标，因为不是按钮 */
  transition: all 0.3s ease;
  position: relative;
}

.setting-option:hover {
  background: linear-gradient(
    90deg,
    rgba(173, 216, 230, 0.6),  /* 粉碎冰蓝色调 */
    rgba(255, 182, 193, 0.9)   /* 玫瑰色调 */
  );
  color: #1e8fa3;
  text-shadow: 0 1px 2px rgba(255, 255, 255, 0.1);
}

.setting-option::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.setting-option:hover::before {
  opacity: 1;
}

.option-label {
  /* 左侧标签样式，靠左对齐 */
  flex: 0 0 auto; /* 不拉伸，占据内容宽度 */
  margin-right: auto; /* 推动控件到右侧 */
  text-align: left;
}

.option-control {
  /* 右侧控件容器 */
  flex: 0 0 auto; /* 不拉伸，占据内容宽度 */
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

/* 开关按钮样式 */
.switch {
  position: relative;
  display: inline-block;
  width: 40px;
  height: 20px;
  margin-left: 10px;
}
.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}
.switch::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  border-radius: 34px;
  transition: 0.4s;
}
.switch::after {
  content: "";
  position: absolute;
  height: 14px;
  width: 14px;
  left: 3px;
  bottom: 3px;
  background-color: white;
  border-radius: 50%;
  transition: 0.4s;
}
.switch input:checked + .switch::before {
  background-color: #4eb157;
}
.switch input:checked + .switch::after {
  transform: translateX(20px);
}

/* 输入框样式 */
.setting-option input[type="text"],
.setting-option textarea {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  background: rgba(255, 255, 255, 0.6);
  transition: background 0.3s ease;
}
.setting-option input[type="text"]:focus,
.setting-option textarea:focus {
  background: rgba(255, 255, 255, 0.1);
  outline: none;
}

/* 暗色模式样式 */
@media (prefers-color-scheme: dark) {
  .setting-option {
    color: rgba(255, 255, 255, 0.9);
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5);
    background: linear-gradient(
      90deg,
      rgba(255, 255, 255, 0.15),
      rgba(255, 255, 255, 0.1)
    );
  }

  .setting-option:hover {
    color: #ffffff;
  }

  .switch::before {
    background-color: rgba(255, 255, 255, 0.3);
  }

  .switch input:checked + .switch::before {
    background-color: #4eb157;
  }

  .setting-option input[type="text"],
  .setting-option textarea {
    background: rgba(255, 255, 255, 0.5);
    color: #000;
  }

  .setting-option input[type="text"]:focus,
  .setting-option textarea:focus {
    background: rgba(255, 255, 255, 0.7);
  }
}
</style>