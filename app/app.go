package app

import (
	xc "github.com/twgh/xcgui"
)

type App struct {
}

// 炫彩_初始化
//  pText: 保留参数.
func New(pText string) *App {
	p := &App{}
	if !xc.XInitXCGUI(pText) {
		return nil
	}
	return p
}

// 炫彩_运行, 运行消息循环,当炫彩窗口数量为0时退出
func (a *App) Run() int {
	return xc.XRunXCGUI()
}

// 炫彩_退出, 退出界面库释放资源
func (a *App) Exit() int {
	return xc.XExitXCGUI()
}

// 炫彩_输出调试信息到文件, 打印调试信息到文件xcgui_debug.txt
func (a *App) DebugToFileInfo(pInfo string) int {
	return xc.XC_DebugToFileInfo(pInfo)
}

// 炫彩_激活窗口, 激活当前进程最上层窗口
func (a *App) SetActivateTopWindow() bool {
	return xc.XC_SetActivateTopWindow()
}

// 炫彩_取默认字体, 返回默认字体句柄
func (a *App) GetDefaultFont() int {
	return xc.XC_GetDefaultFont()
}

// 炫彩_消息框
// hWindow: 窗口句柄
// pText: 内容文本
// pCaption: 标题
// nFlags: 标识, MessageBox_Flag_
func (a *App) MessageBox(hWindow int, pText string, pCaption string, nFlags int) int {
	return xc.XC_MessageBox(hWindow, pText, pCaption, nFlags)
}

// 炫彩_发送窗口消息
// hWindow: 窗口句柄
// msg:
// wParam:
// lParam:
func (a *App) SendMessage(hWindow int, msg int, wParam int, lParam int) int {
	return xc.XC_SendMessage(hWindow, msg, wParam, lParam)
}

// 炫彩_投递窗口消息
// hWindow: 窗口句柄
// msg:
// wParam:
// lParam:
func (a *App) PostMessage(hWindow int, msg int, wParam int, lParam int) bool {
	return xc.XC_PostMessage(hWindow, msg, wParam, lParam)
}

// 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI
// pCall: 回调函数
// data: 用户自定义数据
func (a *App) CallUiThread(pCall func(data int) int, data int) int {
	return xc.XC_CallUiThread(pCall, data)
}

// 炫彩_判断元素, 判断是否为元素句柄
// hEle: 元素句柄.
func (a *App) IsHELE(hEle int) bool {
	return xc.XC_IsHELE(hEle)
}

// 炫彩_判断窗口, 判断是否为窗口句柄
// hWindow: 窗口句柄.
func (a *App) IsHWINDOW(hWindow int) bool {
	return xc.XC_IsHWINDOW(hWindow)
}

// 炫彩_判断形状对象, 判断是否为形状对象
// hShape: 形状对象句柄.
func (a *App) IsShape(hShape int) bool {
	return xc.XC_IsShape(hShape)
}

// 炫彩_判断句柄包含类型, 判断句柄是否拥有该类型
// hXCGUI: 炫彩句柄.
// nType: 句柄类型, XC_OBJECT_TYPE, 以XC_开头的常量
func (a *App) IsHXCGUI(hXCGUI int, nType int) bool {
	return xc.XC_IsHXCGUI(hXCGUI, nType)
}

// 炫彩_转换HWND到HWINDOW, 通过窗口HWND句柄获取HWINDOW句柄
// hWnd: 窗口HWND句柄.
func (a *App) HWindowFromHWnd(hWnd int) int {
	return xc.XC_hWindowFromHWnd(hWnd)
}

// 炫彩_置属性, 设置对象属性
// hXCGUI: 对象句柄.
// pName: 属性名.
// pValue: 属性值.
func (a *App) SetProperty(hXCGUI int, pName string, pValue string) bool {
	return xc.XC_SetProperty(hXCGUI, pName, pValue)
}

// 炫彩_取属性, 获取对象属性, 返回属性值
// hXCGUI: 对象句柄.
// pName: 属性名.
func (a *App) GetProperty(hXCGUI int, pName string) string {
	return xc.XC_GetProperty(hXCGUI, pName)
}

// 炫彩_注册窗口类名, 如果是在DLL中使用, 那么DLL卸载时需要注销窗口类名, 否则DLL卸载后, 类名所指向的窗口过程地址失效
// pClassName: 类名.
func (a *App) RegisterWindowClassName(pClassName string) bool {
	return xc.XC_RegisterWindowClassName(pClassName)
}

// 炫彩_判断滚动视图扩展元素, 判断元素是否从滚动视图元素扩展的新元素, 包含滚动视图元素
// hEle: 元素句柄.
func (a *App) IsSViewExtend(hEle int) bool {
	return xc.XC_IsSViewExtend(hEle)
}

// 炫彩_取对象类型, 获取句柄类型, 返回: XC_OBJECT_TYPE
// hXCGUI: 炫彩对象句柄.
func (a *App) GetObjectType(hXCGUI int) int {
	return xc.XC_GetObjectType(hXCGUI)
}

// 炫彩_取对象从ID, 通过ID获取对象句柄, 不包括窗口对象
// hWindow: 所属窗口句柄
// nID: ID值.
func (a *App) GetObjectByID(hWindow int, nID int) int {
	return xc.XC_GetObjectByID(hWindow, nID)
}

// 炫彩_取对象从ID名称, 通过ID名称获取对象句柄
// hWindow: 所属窗口句柄
// pName: ID名称.
func (a *App) GetObjectByIDName(hWindow int, pName string) int {
	return xc.XC_GetObjectByIDName(hWindow, pName)
}

// 炫彩_取对象从UID, 通过UID获取对象句柄, 不包括窗口对象
// nUID: UID值.
func (a *App) GetObjectByUID(nUID int) int {
	return xc.XC_GetObjectByUID(nUID)
}

// 炫彩_取对象从UID名称, 通过UID名称获取对象句柄
// pName: UID名称.
func (a *App) GetObjectByUIDName(pName string) int {
	return xc.XC_GetObjectByUIDName(pName)
}

// 炫彩_取对象从名称, 通过name获取对象句柄
// pName: name名称.
func (a *App) GetObjectByName(pName string) int {
	return xc.XC_GetObjectByName(pName)
}

// 炫彩_置绘制频率, 设置UI的最小重绘频率
// nMilliseconds: 重绘最小时间间隔, 单位毫秒
func (a *App) SetPaintFrequency(nMilliseconds int) int {
	return xc.XC_SetPaintFrequency(nMilliseconds)
}

// 炫彩_置文本渲染质量, 设置文本渲染质量
// nType: 参见GDI+ TextRenderingHint 定义.
func (a *App) SetTextRenderingHint(nType int) int {
	return xc.XC_SetTextRenderingHint(nType)
}

// 炫彩_启用GDI绘制文本, 将影响到以下函数: XDraw_TextOut, XDraw_TextOutEx, XDraw_TextOutA
// bEnable: 是否启用
func (a *App) EnableGdiDrawText(bEnable bool) int {
	return xc.XC_EnableGdiDrawText(bEnable)
}

// 炫彩_判断矩形相交, 判断两个矩形是否相交及重叠
// pRect1: 矩形1.
// pRect2: 矩形2.
func (a *App) RectInRect(pRect1 *xc.RECT, pRect2 *xc.RECT) bool {
	return xc.XC_RectInRect(pRect1, pRect2)
}

// 炫彩_组合矩形, 组合两个矩形区域
// pDest: 新的矩形区域.
// pSrc1: 源矩形1.
// pSrc2: 源矩形2.
func (a *App) CombineRect(pDest *xc.RECT, pSrc1 *xc.RECT, pSrc2 *xc.RECT) int {
	return xc.XC_CombineRect(pDest, pSrc1, pSrc2)
}

// 炫彩_显示布局边界, 显示布局对象边界
// bShow: 是否显示.
func (a *App) ShowLayoutFrame(bShow bool) int {
	return xc.XC_ShowLayoutFrame(bShow)
}

// 炫彩_启用debug文件
// bEnable: 是否启用.
func (a *App) EnableDebugFile(bEnable bool) int {
	return xc.XC_EnableDebugFile(bEnable)
}

// 炫彩_启用资源监视器
// bEnable: 是否启用.
func (a *App) EnableResMonitor(bEnable bool) int {
	return xc.XC_EnableResMonitor(bEnable)
}

// 炫彩_置布局边界颜色
// color: RGB颜色值.
func (a *App) SetLayoutFrameColor(color int) int {
	return xc.XC_SetLayoutFrameColor(color)
}

// 炫彩_启用错误弹窗, 启用错误弹出, 通过该接口可以设置遇到严重错误时不弹出消息提示框
// bEnabel: 是否启用.
func (a *App) EnableErrorMessageBox(bEnabel bool) int {
	return xc.XC_EnableErrorMessageBox(bEnabel)
}

// 炫彩_启用自动退出程序, 启动或禁用自动退出程序, 当检测到所有用户创建的窗口都关闭时, 自动退出程序; 可调用 XC_PostQuitMessage() 手动退出程序
// bEnabel: 是否启用.
func (a *App) EnableAutoExitApp(bEnabel bool) int {
	return xc.XC_EnableAutoExitApp(bEnabel)
}

// 炫彩_取文本绘制大小
// pString: 字符串.
// length: 字符串长度
// hFontX: 字体.
// pOutSize: 接收返回大小.
func (a *App) GetTextSize(pString string, length int, hFontX int, pOutSize *xc.SIZE) int {
	return xc.XC_GetTextSize(pString, length, hFontX, pOutSize)
}

// 炫彩_取文本显示大小
// pString: 字符串.
// length: 字符串长度
// hFontX: 字体.
// pOutSize: 接收返回大小.
func (a *App) GetTextShowSize(pString string, length int, hFontX int, pOutSize *xc.SIZE) int {
	return xc.XC_GetTextShowSize(pString, length, hFontX, pOutSize)
}

// 炫彩_取文本显示大小扩展
// pString: 字符串.
// length: 字符串长度
// hFontX: 字体.
// nTextAlign: 文本对齐方式, TextFormatFlag_
// pOutSize: 接收返回大小.
func (a *App) GetTextShowSizeEx(pString string, length int, hFontX int, nTextAlign int, pOutSize *xc.SIZE) int {
	return xc.XC_GetTextShowSizeEx(pString, length, hFontX, nTextAlign, pOutSize)
}

// 炫彩_取文本显示矩形
// pString: 字符串.
// length: 字符串长度
// hFontX: 字体.
// width: 最大宽度
// pOutSize: 接收返回大小.
func (a *App) GetTextShowRect(pString string, length int, hFontX int, width int, pOutSize *xc.SIZE) int {
	return xc.XC_GetTextShowRect(pString, length, hFontX, width, pOutSize)
}

// 炫彩_置默认字体
// hFontX: 炫彩字体句柄.
func (a *App) SetDefaultFont(hFontX int) int {
	return xc.XC_SetDefaultFont(hFontX)
}

// 炫彩_添加搜索路径, 添加文件搜索路径, 默认路为exe目录和程序当前运行目录
// pPath: 文件夹
func (a *App) AddFileSearchPath(pPath string) int {
	return xc.XC_AddFileSearchPath(pPath)
}

// 炫彩_初始化字体, 初始化LOGFONTW结构体
// pFont: LOGFONTW结构体指针.
// pName: 字体名称.
// size: 字体大小.
// bBold: 是否为粗体.
// bItalic: 是否为斜体.
// bUnderline: 是否有下划线.
// bStrikeOut: 是否有删除线.
func (a *App) InitFont(pFont *xc.LOGFONTW, pName string, size int, bBold bool, bItalic bool, bUnderline bool, bStrikeOut bool) int {
	return xc.XC_InitFont(pFont, pName, size, bBold, bItalic, bUnderline, bStrikeOut)
}

// 炫彩_分配内存, 在UI库中申请内存, 返回: 内存首地址.
// size: 大小, 字节为单位
func (a *App) Malloc(size int) int {
	return xc.XC_Malloc(size)
}

// 炫彩_释放内存, 在UI库中释放内存
// p: 内存首地址.
func (a *App) Free(p int) int {
	return xc.XC_Free(p)
}

// 炫彩_弹框, 弹出提示框
// pText: 提示内容
// pTitle: 提示框标题
func (a *App) Alert(pText string, pTitle string) int {
	return xc.XC_Alert(pText, pTitle)
}

// 炫彩_系统_ShellExecute, 参见系统API ShellExecute()
// hwnd:
// lpOperation:
// lpFile:
// lpParameters:
// lpDirectory:
// nShowCmd:
func (a *App) Sys_ShellExecute(hwnd int, lpOperation string, lpFile string, lpParameters string, lpDirectory string, nShowCmd int) int {
	return xc.XC_Sys_ShellExecute(hwnd, lpOperation, lpFile, lpParameters, lpDirectory, nShowCmd)
}

// 炫彩_载入动态库, 系统API LoadLibrary, 返回动态库模块句柄
// lpFileName: 文件名
func (a *App) LoadLibrary(lpFileName string) int {
	return xc.XC_LoadLibrary(lpFileName)
}

// 炫彩_取动态库中函数地址, 系统API GetProcAddress, 返回函数地址
// hModule: 动态库模块句柄
// lpProcName: 函数名
func (a *App) GetProcAddress(hModule int, lpProcName string) int {
	return xc.XC_GetProcAddress(hModule, lpProcName)
}

// 炫彩_释放动态库, 系统API FreeLibrary
// hModule: 动态库模块句柄
func (a *App) FreeLibrary(hModule int) bool {
	return xc.XC_FreeLibrary(hModule)
}

// 炫彩_加载DLL, 返回DLL模块句柄. 加载指定DLL, 并且调用DLL中函数LoadDll(), DLL中导出函数格式: int WINAPI LoadDll()
// pDllFileName: DLL文件名
func (a *App) LoadDll(pDllFileName string) int {
	return xc.XC_LoadDll(pDllFileName)
}

// 炫彩_PostQuitMessage, 发送WM_QUIT消息退出消息循环
// nExitCode: 退出码.
func (a *App) PostQuitMessage(nExitCode int) int {
	return xc.XC_PostQuitMessage(nExitCode)
}