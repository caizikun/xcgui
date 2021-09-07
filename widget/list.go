package widget

import (
	"github.com/twgh/xcgui/xc"
)

type List struct {
	Element

	HEle int
}

// 列表_创建, 创建列表元素, 返回元素句柄
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素
func XList_Create(x int, y int, cx int, cy int, hParent int) *List {
	p := &List{
		HEle: xc.XList_Create(x, y, cx, cy, hParent),
	}
	p.HEle_ = p.HEle
	return p
}

// 列表_增加列, 返回位置索引
// width: 列宽度.
func (l *List) AddColumn(width int) int {
	return xc.XList_AddColumn(l.HEle, width)
}

// 列表_插入列, 返回插入位置索引
// width: 列宽度.
// iItem: 插入位置索引.
func (l *List) InsertColumn(width int, iItem int) int {
	return xc.XList_InsertColumn(l.HEle, width, iItem)
}

// 列表_启用多选, 启用或关闭多选功能
// bEnable: 是否启用.
func (l *List) EnableMultiSel(bEnable bool) int {
	return xc.XList_EnableMultiSel(l.HEle, bEnable)
}

// 列表_启用拖动更改列宽, 启用拖动改变列宽度
// bEnable: 是否启用.
func (l *List) EnableDragChangeColumnWidth(bEnable bool) int {
	return xc.XList_EnableDragChangeColumnWidth(l.HEle, bEnable)
}

// 列表_启用垂直滚动条顶部对齐
// bTop: 是否启用.
func (l *List) EnableVScrollBarTop(bTop bool) int {
	return xc.XList_EnableVScrollBarTop(l.HEle, bTop)
}

// 列表_启用项背景全行模式, 启用项背景全行填充模式
// bFull: 是否启用.
func (l *List) EnableItemBkFullRow(bFull bool) int {
	return xc.XList_EnableItemBkFullRow(l.HEle, bFull)
}

// 列表_启用固定行高
// bEnable: 是否启用
func (l *List) EnableFixedRowHeight(bEnable bool) int {
	return xc.XList_EnableFixedRowHeight(l.HEle, bEnable)
}

// 列表_启用模板复用
// bEnable: 是否启用
func (l *List) EnablemTemplateReuse(bEnable bool) int {
	return xc.XList_EnablemTemplateReuse(l.HEle, bEnable)
}

// 列表_启用虚表
// bEnable: 是否启用
func (l *List) EnableVirtualTable(bEnable bool) int {
	return xc.XList_EnableVirtualTable(l.HEle, bEnable)
}

// 列表_置虚表行数
// nRowCount: 行数
func (l *List) SetVirtualRowCount(nRowCount int) int {
	return xc.XList_SetVirtualRowCount(l.HEle, nRowCount)
}

// 列表_置排序, 设置排序属性
// iColumn: 列索引.
// iColumnAdapter: 需要排序的数据在数据适配器中的列索引.
// bEnable: 是否启用排序功能.
func (l *List) SetSort(iColumn int, iColumnAdapter int, bEnable bool) int {
	return xc.XList_SetSort(l.HEle, iColumn, iColumnAdapter, bEnable)
}

// 列表_置绘制项背景标志, 设置是否绘制指定状态下项的背景
// nFlags: 标志位, List_DrawItemBk_Flag_
func (l *List) SetDrawItemBkFlags(nFlags int) int {
	return xc.XList_SetDrawItemBkFlags(l.HEle, nFlags)
}

// 列表_置列宽
// iItem: 列索引.
// width: 宽度.
func (l *List) SetColumnWidth(iItem int, width int) int {
	return xc.XList_SetColumnWidth(l.HEle, iItem, width)
}

// 列表_置列最小宽度
// iItem: 列索引.
// width: 宽度.
func (l *List) SetColumnMinWidth(iItem int, width int) int {
	return xc.XList_SetColumnMinWidth(l.HEle, iItem, width)
}

// 列表_置列宽度固定
// iColumn: 列索引.
// bFixed: 是否固定宽度.
func (l *List) SetColumnWidthFixed(iColumn int, bFixed bool) int {
	return xc.XList_SetColumnWidthFixed(l.HEle, iColumn, bFixed)
}

// 列表_取列宽度
// iColumn: 列索引.
func (l *List) GetColumnWidth(iColumn int) int {
	return xc.XList_GetColumnWidth(l.HEle, iColumn)
}

// 列表_取列数量
func (l *List) GetColumnCount() int {
	return xc.XList_GetColumnCount(l.HEle)
}

// 列表_置项数据, 设置项用户数据
// iItem: 项索引.
// iSubItem: 子项索引.
// data: 用户数据.
func (l *List) SetItemData(iItem int, iSubItem int, data int) bool {
	return xc.XList_SetItemData(l.HEle, iItem, iSubItem, data)
}

// 列表_取项数据, 获取项用户数据
// iItem: 项索引.
// iSubItem: 子项索引.
func (l *List) GetItemData(iItem int, iSubItem int) int {
	return xc.XList_GetItemData(l.HEle, iItem, iSubItem)
}

// 列表_置选择项
// iItem: 项索引.
func (l *List) SetSelectItem(iItem int) bool {
	return xc.XList_SetSelectItem(l.HEle, iItem)
}

// 列表_取选择项, 返回项索引
func (l *List) GetSelectItem() int {
	return xc.XList_GetSelectItem(l.HEle)
}

// 列表_取选择项数量, 获取选择项数量
func (l *List) GetSelectItemCount() int {
	return xc.XList_GetSelectItemCount(l.HEle)
}

// 列表_添加选择项
// iItem: 项索引
func (l *List) AddSelectItem(iItem int) bool {
	return xc.XList_AddSelectItem(l.HEle, iItem)
}

// 列表_置选择全部, 选择全部行
func (l *List) SetSelectAll() int {
	return xc.XList_SetSelectAll(l.HEle)
}

// 列表_取全部选择, 获取全部选择的行, 返回行数量
// pArray: 接收行索引数组.
// nArraySize: 数组大小.
func (l *List) GetSelectAll(pArray int, nArraySize int) int {
	return xc.XList_GetSelectAll(l.HEle, pArray, nArraySize)
}

// 列表_显示指定项, 滚动视图让指定项可见
// iItem: 项索引.
func (l *List) VisibleItem(iItem int) int {
	return xc.XList_VisibleItem(l.HEle, iItem)
}

// 列表_取消选择项, 取消选择指定项(这里的项可以理解为行)
// iItem: 项索引.
func (l *List) CancelSelectItem(iItem int) bool {
	return xc.XList_CancelSelectItem(l.HEle, iItem)
}

// 列表_取消全部选择项, 取消选择所有项(这里的项可以理解为行)
func (l *List) CancelSelectAll() int {
	return xc.XList_CancelSelectAll(l.HEle)
}

// 列表_取列表头, 获取列表头元素, 返回列表头元素句柄
func (l *List) GetHeaderHELE() int {
	return xc.XList_GetHeaderHELE(l.HEle)
}

// 列表_删除列
// iItem: 项索引.
func (l *List) DeleteColumn(iItem int) bool {
	return xc.XList_DeleteColumn(l.HEle, iItem)
}

// 列表_删除全部列, 删除所有的列, 仅删除List的, 数据适配器的列不变
func (l *List) DeleteColumnAll() int {
	return xc.XList_DeleteColumnAll(l.HEle)
}

// 列表_绑定数据适配器
// hAdapter: 数据适配器句柄 XAdTable.
func (l *List) BindAdapter(hAdapter int) int {
	return xc.XList_BindAdapter(l.HEle, hAdapter)
}

// 列表_列表头绑定数据适配器
// hAdapter: 数据适配器句柄 XAdMap.
func (l *List) BindAdapterHeader(hAdapter int) int {
	return xc.XList_BindAdapterHeader(l.HEle, hAdapter)
}

// 列表_创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (l *List) CreateAdapter() int {
	return xc.XList_CreateAdapter(l.HEle)
}

// 列表_列表头创建数据适配器, 创建数据适配器，根据绑定的项模板初始化数据适配器的列, 返回适配器句柄.
func (l *List) CreateAdapterHeader() int {
	return xc.XList_CreateAdapterHeader(l.HEle)
}

// 列表_取数据适配器, 返回数据适配器句柄
func (l *List) GetAdapter() int {
	return xc.XList_GetAdapter(l.HEle)
}

// 列表_列表头获取数据适配器, 获取列表头数据适配器句柄
func (l *List) GetAdapterHeader() int {
	return xc.XList_GetAdapterHeader(l.HEle)
}

// 列表_置项模板文件, 设置项布局模板文件
// pXmlFile: 文件名.
func (l *List) SetItemTemplateXML(pXmlFile string) bool {
	return xc.XList_SetItemTemplateXML(l.HEle, pXmlFile)
}

// 列表_置项模板从字符串, 设置项布局模板文件
// pStringXML: 字符串指针.
func (l *List) SetItemTemplateXMLFromString(pStringXML int) bool {
	return xc.XList_SetItemTemplateXMLFromString(l.HEle, pStringXML)
}

// 列表_置项模板, 设置列表项模板
// hTemp: 模板句柄.
func (l *List) SetItemTemplate(hTemp int) bool {
	return xc.XList_SetItemTemplate(l.HEle, hTemp)
}

// 列表_取项模板对象, 通过模板项ID, 获取实例化模板项ID对应的对象句柄
// iItem: 项索引.
// iSubItem: 子项索引.
// nTempItemID: 模板项itemID.
func (l *List) GetTemplateObject(iItem int, iSubItem int, nTempItemID int) int {
	return xc.XList_GetTemplateObject(l.HEle, iItem, iSubItem, nTempItemID)
}

// 列表_取对象所在行, 获取当前对象所在模板实例, 属于列表中哪一个项. 成功返回项索引, 否则返回XC_ID_ERROR.
// hXCGUI: 对象句柄, UI元素句柄或形状对象句柄
func (l *List) GetItemIndexFromHXCGUI(hXCGUI int) int {
	return xc.XList_GetItemIndexFromHXCGUI(l.HEle, hXCGUI)
}

// 列表_取列表头模板对象, 列表头, 通过模板项ID, 获取实例化模板项ID对应的对象句柄
// iItem: 列表头项ID.
// nTempItemID: 模板项ID.
func (l *List) GetHeaderTemplateObject(iItem int, nTempItemID int) int {
	return xc.XList_GetHeaderTemplateObject(l.HEle, iItem, nTempItemID)
}

// 列表_取列表头对象所在行, 列表头, 获取当前对象所在模板实例, 属于列表头中哪一个项. 成功返回项索引, 否则返回XC_ID_ERROR.
// hXCGUI: 对象句柄.
func (l *List) GetHeaderItemIndexFromHXCGUI(hXCGUI int) int {
	return xc.XList_GetHeaderItemIndexFromHXCGUI(l.HEle, hXCGUI)
}

// 列表_置列表头高度
// height: 高度.
func (l *List) SetHeaderHeight(height int) int {
	return xc.XList_SetHeaderHeight(l.HEle, height)
}

// 列表_取列表头高度
func (l *List) GetHeaderHeight() int {
	return xc.XList_GetHeaderHeight(l.HEle)
}

// 列表_取可视行范围
// piStart: 开始行索引.
// piEnd: 结束行索引.
func (l *List) GetVisibleRowRange(piStart *int, piEnd *int) int {
	return xc.XList_GetVisibleRowRange(l.HEle, piStart, piEnd)
}

// 列表_添加项背景边框, 添加项背景内容边框
// nState: 项状态.
// color: RGB颜色.
// alpha: 透明度.
// width: 线宽.
func (l *List) AddItemBkBorder(nState int, color int, alpha uint8, width int) int {
	return xc.XList_AddItemBkBorder(l.HEle, nState, color, alpha, width)
}

// 列表_添加项背景填充, 添加项背景内容填充
// nState: 项状态.
// color: RGB颜色.
// alpha: 透明度.
func (l *List) AddItemBkFill(nState int, color int, alpha uint8) int {
	return xc.XList_AddItemBkFill(l.HEle, nState, color, alpha)
}

// 列表_添加项背景图片, 添加项背景内容图片
// nState: 项状态.
// hImage: 图片句柄.
func (l *List) AddItemBkImage(nState int, hImage int) int {
	return xc.XList_AddItemBkImage(l.HEle, nState, hImage)
}

// 列表_取项背景对象数量, 成功返回背景内容数量, 否则返回XC_ID_ERROR
func (l *List) GetItemBkInfoCount() int {
	return xc.XList_GetItemBkInfoCount(l.HEle)
}

// 列表_清空项背景对象, 清空项背景内容; 如果背景没有内容, 将使用系统默认内容, 以便保证背景正确.
func (l *List) ClearItemBkInfo() int {
	return xc.XList_ClearItemBkInfo(l.HEle)
}

// 列表_置项默认高度
// nHeight: 高度.
// nSelHeight: 选中时高度.
func (l *List) SetItemHeightDefault(nHeight int, nSelHeight int) int {
	return xc.XList_SetItemHeightDefault(l.HEle, nHeight, nSelHeight)
}

// 列表_取项默认高度
// pHeight: 高度.
// pSelHeight: 选中时高度.
func (l *List) GetItemHeightDefault(pHeight *int, pSelHeight *int) int {
	return xc.XList_GetItemHeightDefault(l.HEle, pHeight, pSelHeight)
}

// 列表_置行间距
// nSpace: 行间距大小.
func (l *List) SetRowSpace(nSpace int) int {
	return xc.XList_SetRowSpace(l.HEle, nSpace)
}

// 列表_取行间距
func (l *List) GetRowSpace() int {
	return xc.XList_GetRowSpace(l.HEle)
}

// 列表_置锁定列左侧, 锁定列, 设置左侧锁定列分界列索引
// iColumn: 列索引, -1代表不锁定
func (l *List) SetLockColumnLeft(iColumn int) int {
	return xc.XList_SetLockColumnLeft(l.HEle, iColumn)
}

// 列表_置锁定列右侧
// iColumn: 列索引, -1代表不锁定. 暂时只支持锁定末尾列
func (l *List) SetLockColumnRight(iColumn int) int {
	return xc.XList_SetLockColumnRight(l.HEle, iColumn)
}

// 列表_置锁定行底部, 设置是否锁定末尾行
// bLock: 是否锁定.
func (l *List) SetLockRowBottom(bLock bool) int {
	return xc.XList_SetLockRowBottom(l.HEle, bLock)
}

// 列表_置锁定行底部重叠
// bOverlap: 是否重叠
func (l *List) SetLockRowBottomOverlap(bOverlap bool) int {
	return xc.XList_SetLockRowBottomOverlap(l.HEle, bOverlap)
}

// 列表_测试点击项, 检测坐标点所在项
// pPt: 坐标点.
// piItem: 项索引.
// piSubItem: 子项索引.
func (l *List) HitTest(pPt *xc.POINT, piItem *int, piSubItem *int) bool {
	return xc.XList_HitTest(l.HEle, pPt, piItem, piSubItem)
}

// 列表_测试点击项扩展, 检查坐标点所在项, 自动添加滚动视图偏移量
// pPt: 坐标点.
// piItem: 项索引.
// piSubItem: 子项索引.
func (l *List) HitTestOffset(pPt *xc.POINT, piItem *int, piSubItem *int) bool {
	return xc.XList_HitTestOffset(l.HEle, pPt, piItem, piSubItem)
}

// 列表_刷新项数据
func (l *List) RefreshData() int {
	return xc.XList_RefreshData(l.HEle)
}

// 列表_刷新指定项, 刷新指定项模板, 以便更新UI
// iItem: 项索引.
func (l *List) RefreshItem(iItem int) int {
	return xc.XList_RefreshItem(l.HEle, iItem)
}

// 列表_添加列文本
// nWidth:
// pName:
// pText:
func (l *List) AddColumnText(nWidth int, pName string, pText string) int {
	return xc.XList_AddColumnText(l.HEle, nWidth, pName, pText)
}

// 列表_添加列图片
// nWidth:
// pName:
// hImage:
func (l *List) AddColumnImage(nWidth int, pName string, hImage int) int {
	return xc.XList_AddColumnImage(l.HEle, nWidth, pName, hImage)
}

// 列表_添加项文本
// pText:
func (l *List) AddItemText(pText string) int {
	return xc.XList_AddItemText(l.HEle, pText)
}

// 列表_添加项文本扩展
// pName:
// pText:
func (l *List) AddItemTextEx(pName string, pText string) int {
	return xc.XList_AddItemTextEx(l.HEle, pName, pText)
}

// 列表_添加项图片
// hImage:
func (l *List) AddItemImage(hImage int) int {
	return xc.XList_AddItemImage(l.HEle, hImage)
}

// 列表_添加项图片扩展
// pName:
// hImage:
func (l *List) AddItemImageEx(pName string, hImage int) int {
	return xc.XList_AddItemImageEx(l.HEle, pName, hImage)
}

// 列表_插入项文本
// iItem:
// pValue:
func (l *List) InsertItemText(iItem int, pValue string) int {
	return xc.XList_InsertItemText(l.HEle, iItem, pValue)
}

// 列表_插入项文本扩展
// iItem:
// pName:
// pValue:
func (l *List) InsertItemTextEx(iItem int, pName string, pValue string) int {
	return xc.XList_InsertItemTextEx(l.HEle, iItem, pName, pValue)
}

// 列表_插入项图片
// iItem:
// hImage:
func (l *List) InsertItemImage(iItem int, hImage int) int {
	return xc.XList_InsertItemImage(l.HEle, iItem, hImage)
}

// 列表_插入项图片扩展
// iItem:
// pName:
// hImage:
func (l *List) InsertItemImageEx(iItem int, pName string, hImage int) int {
	return xc.XList_InsertItemImageEx(l.HEle, iItem, pName, hImage)
}

// 列表_置项文本
// iItem:
// iColumn:
// pText:
func (l *List) SetItemText(iItem int, iColumn int, pText string) bool {
	return xc.XList_SetItemText(l.HEle, iItem, iColumn, pText)
}

// 列表_置项文本扩展
// iItem:
// pName:
// pText:
func (l *List) SetItemTextEx(iItem int, pName string, pText string) bool {
	return xc.XList_SetItemTextEx(l.HEle, iItem, pName, pText)
}

// 列表_置项图片
// iItem:
// iColumn:
// hImage:
func (l *List) SetItemImage(iItem int, iColumn int, hImage int) bool {
	return xc.XList_SetItemImage(l.HEle, iItem, iColumn, hImage)
}

// 列表_置项图片扩展
// iItem:
// pName:
// hImage:
func (l *List) SetItemImageEx(iItem int, pName string, hImage int) bool {
	return xc.XList_SetItemImageEx(l.HEle, iItem, pName, hImage)
}

// 列表_置项指数值
// iItem:
// iColumn:
// nValue:
func (l *List) SetItemInt(iItem int, iColumn int, nValue int) bool {
	return xc.XList_SetItemInt(l.HEle, iItem, iColumn, nValue)
}

// 列表_置项整数值扩展
// iItem:
// pName:
// nValue:
func (l *List) SetItemIntEx(iItem int, pName string, nValue int) bool {
	return xc.XList_SetItemIntEx(l.HEle, iItem, pName, nValue)
}

// 列表_置项浮点值
// iItem:
// iColumn:
// fFloat:
func (l *List) SetItemFloat(iItem int, iColumn int, fFloat float32) bool {
	return xc.XList_SetItemFloat(l.HEle, iItem, iColumn, fFloat)
}

// 列表_置项浮点值扩展
// iItem:
// pName:
// fFloat:
func (l *List) SetItemFloatEx(iItem int, pName string, fFloat float32) bool {
	return xc.XList_SetItemFloatEx(l.HEle, iItem, pName, fFloat)
}

// 列表_取项文本
// iItem:
// iColumn:
func (l *List) GetItemText(iItem int, iColumn int) string {
	return xc.XList_GetItemText(l.HEle, iItem, iColumn)
}

// 列表_取项文本扩展
// iItem:
// pName:
func (l *List) GetItemTextEx(iItem int, pName string) string {
	return xc.XList_GetItemTextEx(l.HEle, iItem, pName)
}

// 列表_取项图片
// iItem:
// iColumn:
func (l *List) GetItemImage(iItem int, iColumn int) int {
	return xc.XList_GetItemImage(l.HEle, iItem, iColumn)
}

// 列表_取项图片扩展
// iItem:
// pName:
func (l *List) GetItemImageEx(iItem int, pName string) int {
	return xc.XList_GetItemImageEx(l.HEle, iItem, pName)
}

// 列表_取项整数值
// iItem:
// iColumn:
// pOutValue:
func (l *List) GetItemInt(iItem int, iColumn int, pOutValue *int) bool {
	return xc.XList_GetItemInt(l.HEle, iItem, iColumn, pOutValue)
}

// 列表_取项整数值扩展
// iItem:
// pName:
// pOutValue:
func (l *List) GetItemIntEx(iItem int, pName string, pOutValue *int) bool {
	return xc.XList_GetItemIntEx(l.HEle, iItem, pName, pOutValue)
}

// 列表_取项浮点值
// iItem:
// iColumn:
// pOutValue:
func (l *List) GetItemFloat(iItem int, iColumn int, pOutValue *float32) bool {
	return xc.XList_GetItemFloat(l.HEle, iItem, iColumn, pOutValue)
}

// 列表_取项浮点值扩展
// iItem:
// pName:
// pOutValue:
func (l *List) GetItemFloatEx(iItem int, pName string, pOutValue *float32) bool {
	return xc.XList_GetItemFloatEx(l.HEle, iItem, pName, pOutValue)
}

// 列表_删除项
// iItem:
func (l *List) DeleteItem(iItem int) bool {
	return xc.XList_DeleteItem(l.HEle, iItem)
}

// 列表_删除项扩展
// iItem:
// nCount:
func (l *List) DeleteItemEx(iItem int, nCount int) bool {
	return xc.XList_DeleteItemEx(l.HEle, iItem, nCount)
}

// 列表_删除项全部
func (l *List) DeleteItemAll() int {
	return xc.XList_DeleteItemAll(l.HEle)
}

// 列表_删除列全部AD
func (l *List) DeleteColumnAll_AD() int {
	return xc.XList_DeleteColumnAll_AD(l.HEle)
}

// 列表_取项数量AD
func (l *List) GetCount_AD() int {
	return xc.XList_GetCount_AD(l.HEle)
}

// 列表_取列数量AD
func (l *List) GetCountColumn_AD() int {
	return xc.XList_GetCountColumn_AD(l.HEle)
}