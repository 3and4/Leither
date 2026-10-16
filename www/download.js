// 公共基础URL
const BASE_PATH = "http://127.0.0.1:4800/mm/Fc1BRTFafOGzq5P8KmkVJqwS2v2";
// 版本号
const VERSION = "V0.21.06";
// 完整基础URL
const BASE_URL = `${BASE_PATH}/${VERSION}`;

/**
 * 架构映射表
 * 将文件名后缀映射为可读的架构名称
 */
const ARCH_MAP = {
  'amd64': 'x86-64',
  'arm64': 'ARM64',
  'arm7': 'ARMv7'
};

/**
 * 操作系统映射表
 * 将文件名中的系统标识映射为可读的系统名称
 */
const OS_MAP = {
  'darwin': 'macOS',
  'linux': 'Linux',
  'windows': 'Windows'
};

/**
 * 下载数据配置
 * 现在只需要文件名和文件大小
 */
const downloadData = [
  {
    fileName: "Leither.darwin.amd64",
    size: "21.9M"
  },
  {
    fileName: "Leither.darwin.arm64",
    size: "61.9M"
  },
  {
    fileName: "Leither.linux.amd64",
    size: "21M"
  },
  {
    fileName: "Leither.linux.arm64",
    size: "18.7M"
  },
  {
    fileName: "Leither.linux.arm7",
    size: "18M"
  },
  {
    fileName: "Leither.windows.amd64",
    size: "22M"
  }
];

/**
 * 从文件名解析操作系统信息
 * @param {string} fileName - 完整文件名
 * @returns {string} 操作系统名称
 */
function parseOS(fileName) {
  const parts = fileName.split('.');
  const osKey = parts[parts.length - 2]; // 获取倒数第二部分作为系统key
  return OS_MAP[osKey] || osKey; // 返回映射值或原值
}

/**
 * 从文件名解析架构信息
 * @param {string} fileName - 完整文件名
 * @returns {string} 架构名称
 */
function parseArch(fileName) {
  const parts = fileName.split('.');
  const archKey = parts[parts.length - 1]; // 获取最后部分作为架构key
  return ARCH_MAP[archKey] || archKey; // 返回映射值或原值
}

/**
 * 渲染下载表格
 * 根据downloadData动态生成表格行
 */
function renderDownloadTable() {
  const tbody = document.getElementById('download-table');
  
  // 生成表格行HTML，动态解析操作系统和架构
  const rows = downloadData.map(item => 
    `<tr>
      <td><a href="${BASE_URL}/${item.fileName}">${item.fileName}</a></td>
      <td>${parseOS(item.fileName)}</td>
      <td>${parseArch(item.fileName)}</td>
      <td>${item.size}</td>
    </tr>`
  ).join('');
  
  tbody.innerHTML = rows;
}

// 页面加载完成后初始化表格
window.addEventListener('DOMContentLoaded', renderDownloadTable);