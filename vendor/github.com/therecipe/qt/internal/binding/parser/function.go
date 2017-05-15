package parser

import (
	"fmt"
	"sort"
	"strings"
)

type Function struct {
	Name            string       `xml:"name,attr"`
	Fullname        string       `xml:"fullname,attr"`
	Href            string       `xml:"href,attr"`
	Status          string       `xml:"status,attr"`
	Access          string       `xml:"access,attr"`
	Filepath        string       `xml:"filepath,attr"`
	Virtual         string       `xml:"virtual,attr"`
	Meta            string       `xml:"meta,attr"`
	Static          bool         `xml:"static,attr"`
	Overload        bool         `xml:"overload,attr"`
	OverloadNumber  string       `xml:"overload-number,attr"`
	Output          string       `xml:"type,attr"`
	Signature       string       `xml:"signature,attr"`
	Parameters      []*Parameter `xml:"parameter"`
	Brief           string       `xml:"brief,attr"`
	Since           string       `xml:"since,attr"`
	SignalMode      string
	TemplateModeJNI string
	Default         bool
	TmpName         string
	Export          bool
	NeedsFinalizer  bool
	Container       string
	TemplateModeGo  string
	NonMember       bool
	NoMocDeduce     bool
	AsError         bool
	Synthetic       bool
	Checked         bool
	Exception       bool
	IsMap           bool
	OgParameters    []Parameter
	IsMocFunction   bool
}

type Parameter struct {
	Name    string `xml:"name,attr"`
	Value   string `xml:"left,attr"`
	Right   string `xml:"right,attr"`
	Default string `xml:"default,attr"`
}

func (f *Function) Class() (*Class, bool) {
	var class, ok = State.ClassMap[f.ClassName()]
	return class, ok
}

func (f *Function) ClassName() string {
	var s = strings.Split(f.Fullname, "::")
	if len(s) == 3 {
		return s[1]
	}
	return s[0]
}

func (f *Function) register(m string) {

	if c, ok := f.Class(); !ok {
		State.ClassMap[f.ClassName()] = &Class{
			Name:      f.ClassName(),
			Status:    "commendable",
			Module:    m,
			Access:    "public",
			Functions: []*Function{f},
		}
	} else {
		c.Functions = append(c.Functions, f)
	}
}

//TODO: multipoly [][]string
//TODO: connect/disconnect slot functions + add necessary SIGNAL_* functions (check first if really needed)
func (f *Function) PossiblePolymorphicDerivations(self bool) ([]string, string) {
	var out = make([]string, 0)

	var params = func() []*Parameter {
		if self {
			return []*Parameter{{Name: "ptr", Value: f.ClassName()}}
		}
		return f.Parameters
	}()

	for _, p := range params {
		var c, ok = State.ClassMap[CleanValue(p.Value)]
		if !ok {
			continue
		}

		if f.Meta == CONSTRUCTOR {
			for _, class := range State.ClassMap {
				//TODO: use target to block certain classes
				if ShouldBuildForTarget(strings.TrimPrefix(class.Module, "Qt"), State.Target) &&
					!(class.Name == "QCameraViewfinder" || class.Name == "QGraphicsVideoItem" ||
						class.Name == "QVideoWidget" || class.Name == "QVideoWidgetControl") {
					if class.IsPolymorphic() && class.IsSubClassOf(c.Name) && class.IsSupported() {
						out = append(out, class.Name)
					}
				}
			}
		} else {
			var fc, _ = f.Class()
			for _, class := range SortedClassesForModule(fc.Module, false) {
				if class.IsPolymorphic() && class.IsSubClassOf(c.Name) && class.IsSupported() {
					out = append(out, class.Name)
				}
			}
		}

		//TODO: multipoly
		if len(out) > 0 {
			sort.Stable(sort.StringSlice(out))
			out = append(out, c.Name)
			return out, CleanName(p.Name, p.Value)
		}
	}

	return out, ""
}

func (f *Function) PossibleDerivationsReversedAndRemovedPure(self bool) ([]string, string) {
	if self {
		var fc, ok = f.Class()
		if !ok {
			return make([]string, 0), ""
		}
		var derv = fc.GetAllDerivationsInSameModule()
		var out = make([]string, 0)
		for i := len(derv); i > 0; i-- {
			if !(derv[i-1] == "QAbstractButton" && f.Name == "paintEvent") {
				if c, ok := State.ClassMap[derv[i-1]]; ok && c.IsSupported() {
					out = append(out, derv[i-1])
				}
			}
		}
		out = append(out, fc.Name)
		return out, ""
	}

	var out = make([]string, 0)

	var params = func() []*Parameter {
		if self {
			return []*Parameter{{Name: "ptr", Value: f.ClassName()}}
		}
		return f.Parameters
	}()

	var fc, _ = f.Class()

	for _, p := range params {
		var c, ok = State.ClassMap[CleanValue(p.Value)]
		if !ok {
			continue
		}

		for _, class := range SortedClassesForModule(fc.Module, false) {
			if class.IsSubClassOf(c.Name) && class.IsSupported() &&
				!(class.Name == "QAbstractButton" && f.Name == "paintEvent") {
				out = append(out, class.Name)
			}
		}

		//TODO: multipoly
		if len(out) > 0 {
			sort.Stable(sort.StringSlice(out))
			out = append(out, c.Name)
			return out, CleanName(p.Name, p.Value)
		}
	}

	return out, ""
}

func (f *Function) PossibleDerivationsInAllModules(self bool) ([]string, string) {
	var out = make([]string, 0)

	var params = func() []*Parameter {
		if self {
			return []*Parameter{{Name: "ptr", Value: f.ClassName()}}
		}
		return f.Parameters
	}()

	for _, p := range params {
		var c, ok = State.ClassMap[CleanValue(p.Value)]
		if !ok {
			continue
		}

		for _, class := range State.ClassMap {
			if class.IsSubClassOf(c.Name) && class.IsSupported() &&
				!(class.Name == "QAbstractButton" && f.Name == "paintEvent") {
				out = append(out, class.Name)
			}
		}

		//TODO: multipoly
		if len(out) > 0 {
			sort.Stable(sort.StringSlice(out))
			out = append(out, c.Name)
			return out, CleanName(p.Name, p.Value)
		}
	}

	return out, ""
}

func (f *Function) IsJNIGeneric() bool {

	if f.ClassName() == "QAndroidJniObject" {
		switch f.Name {
		case
			"callMethod",
			"callStaticMethod",

			"getField",
			//"setField", -> uses interface{} if not generic

			"getStaticField",
			//"setStaticField", -> uses interface{} if not generic

			"getObjectField",

			"getStaticObjectField",

			"callObjectMethod",
			"callStaticObjectMethod":
			{
				return true
			}

		case "setStaticField":
			{
				if f.OverloadNumber == "2" || f.OverloadNumber == "4" {
					return true
				}
			}
		}
	}

	return false
}

func (f *Function) IsSupported() bool {

	//if utils.QT_VERSION() == "5.8.0" {
	if f.Fullname == "QJSEngine::newQMetaObject" && f.OverloadNumber == "2" ||
		f.Fullname == "QScxmlTableData::instructions" || f.Fullname == "QScxmlTableData::dataNames" ||
		f.Fullname == "QScxmlTableData::stateMachineTable" ||
		f.Fullname == "QTextToSpeech::voiceChanged" {
		if !strings.Contains(f.Access, "unsupported") {
			f.Access = "unsupported_isBlockedFunction"
		}
		return false
	}
	//}

	switch {
	case
		(f.ClassName() == "QAccessibleObject" || f.ClassName() == "QAccessibleInterface" || f.ClassName() == "QAccessibleWidget" || //QAccessible::State -> quint64
			f.ClassName() == "QAccessibleStateChangeEvent") && (f.Name == "state" || f.Name == "changedStates" || f.Name == "m_changedStates" || f.Name == "setM_changedStates" || f.Meta == CONSTRUCTOR),

		f.Fullname == "QPixmapCache::find" && f.OverloadNumber == "4", //Qt::Key -> int
		(f.Fullname == "QPixmapCache::remove" || f.Fullname == "QPixmapCache::insert") && f.OverloadNumber == "2",
		f.Fullname == "QPixmapCache::replace",

		f.Fullname == "QNdefFilter::appendRecord" && !f.Overload, //QNdefRecord::TypeNameFormat -> uint

		f.ClassName() == "QSimpleXmlNodeModel" && f.Meta == CONSTRUCTOR,

		f.Fullname == "QSGMaterialShader::attributeNames",

		f.ClassName() == "QVariant" && (f.Name == "value" || f.Name == "canConvert"), //needs template

		f.Fullname == "QNdefRecord::isRecordType", f.Fullname == "QScriptEngine::scriptValueFromQMetaObject", //needs template
		f.Fullname == "QScriptEngine::fromScriptValue", f.Fullname == "QJSEngine::fromScriptValue",

		f.ClassName() == "QMetaType" && //needs template
			(f.Name == "hasRegisteredComparators" || f.Name == "registerComparators" ||
				f.Name == "hasRegisteredConverterFunction" || f.Name == "registerConverter" ||
				f.Name == "registerEqualsComparator"),

		State.ClassMap[f.ClassName()].Module == MOC && f.Name == "metaObject", //needed for qtmoc

		f.Fullname == "QSignalBlocker::QSignalBlocker" && f.OverloadNumber == "3", //undefined symbol

		(State.ClassMap[f.ClassName()].IsSubClassOf("QCoreApplication") ||
			f.ClassName() == "QAudioInput" || f.ClassName() == "QAudioOutput") && f.Name == "notify", //redeclared (name collision with QObject)

		f.Fullname == "QGraphicsItem::isBlockedByModalPanel", //** problem

		f.Name == "surfaceHandle", //QQuickWindow && QQuickView //unsupported_cppType(QPlatformSurface)

		f.Name == "QDesignerFormWindowInterface" || f.Name == "QDesignerFormWindowManagerInterface" || f.Name == "QDesignerWidgetBoxInterface", //unimplemented virtual

		f.Fullname == "QNdefNfcSmartPosterRecord::titleRecords", //T<T> output with unsupported output for *_atList
		f.Fullname == "QHelpEngineCore::filterAttributeSets", f.Fullname == "QHelpSearchEngine::query", f.Fullname == "QHelpSearchQueryWidget::query",
		f.Fullname == "QPluginLoader::staticPlugins", f.Fullname == "QSslConfiguration::ellipticCurves", f.Fullname == "QSslConfiguration::supportedEllipticCurves",
		f.Fullname == "QTextFormat::lengthVectorProperty", f.Fullname == "QTextTableFormat::columnWidthConstraints", f.Fullname == "QHelpContentWidget::selectedIndexes",

		f.Fullname == "QListView::indexesMoved", f.Fullname == "QAudioInputSelectorControl::availableInputs", f.Fullname == "QScxmlStateMachine::initialValuesChanged",
		f.Fullname == "QAudioOutputSelectorControl::availableOutputs", f.Fullname == "QQuickWebEngineProfile::downloadFinished",
		f.Fullname == "QQuickWindow::closing", f.Fullname == "QQuickWebEngineProfile::downloadRequested", f.Fullname == "QWebEnginePage::fullScreenRequested",

		f.Fullname == "QApplication::autoMaximizeThreshold", f.Fullname == "QApplication::setAutoMaximizeThreshold",

		f.Fullname == "QWebPluginFactory::__plugins_newList",

		strings.Contains(f.Access, "unsupported"):
		{
			if !strings.Contains(f.Access, "unsupported") {
				f.Access = "unsupported_isBlockedFunction"
			}
			return false
		}
	}

	if f.Name == "__draw_selections_newList" {
		return false
	}

	//generic blocked
	//TODO: also check _setList _atList _newList _keyList instead ?
	var genName = strings.TrimPrefix(f.Name, "__")
	if strings.HasPrefix(genName, "registeredTimers") || strings.HasPrefix(genName, "countriesForLanguage") ||
		strings.HasPrefix(genName, "writingSystem") || strings.HasPrefix(genName, "textList") ||
		strings.HasPrefix(genName, "attributes") || strings.HasPrefix(genName, "additionalFormats") ||
		strings.HasPrefix(genName, "rawHeaderPairs") || strings.HasPrefix(genName, "tabs") ||
		strings.HasPrefix(genName, "QInputMethodEvent_attributes") || strings.HasPrefix(genName, "selections") || strings.HasPrefix(genName, "setSelections") ||
		strings.HasPrefix(genName, "formats") || strings.HasPrefix(genName, "setAdditionalFormats") || strings.HasPrefix(genName, "setFormats") ||
		strings.HasPrefix(genName, "setTabs") || strings.HasPrefix(genName, "extraSelections") ||
		strings.HasPrefix(genName, "setExtraSelections") || strings.HasPrefix(genName, "setButtonLayout") ||
		strings.HasPrefix(genName, "setWhiteList") || strings.HasPrefix(genName, "whiteList") ||
		strings.HasPrefix(genName, "supportedViewfinderFrameRateRanges") || strings.HasPrefix(genName, "hits") ||
		strings.HasPrefix(genName, "featureTypes") || strings.HasPrefix(genName, "supportedPaperSources") ||
		strings.HasPrefix(genName, "setTextureData") || strings.HasPrefix(genName, "textureData") ||
		strings.HasPrefix(genName, "QCustom3DVolume_textureData") || strings.HasPrefix(genName, "createTextureData") ||
		strings.Contains(genName, "alternateSubjectNames") || strings.HasPrefix(genName, "fromVariantMap") ||
		strings.HasPrefix(genName, "QScxmlDataModel") {
		return false
	}

	//TODO: blocked for small
	if f.Fullname == "QTemporaryFile::open" && f.OverloadNumber == "2" ||
		f.Fullname == "QXmlEntityResolver::resolveEntity" ||
		f.Fullname == "QXmlReader::parse" && f.OverloadNumber == "2" ||
		f.Fullname == "QGraphicsItem::updateMicroFocus" ||
		f.Fullname == "QGridLayout::addItem" && f.OverloadNumber == "2" ||
		f.Fullname == "QSvgGenerator::metric" ||
		f.Fullname == "QScxmlDataModel::setScxmlEvent" ||
		f.Fullname == "QPageSetupDialog::open" ||
		f.Fullname == "QPrintPreviewDialog::open" ||
		f.Fullname == "QSqlRelationalTableModel::revert" ||
		f.Fullname == "QSqlRelationalTableModel::submit" ||
		f.Fullname == "QSqlTableModel::revert" ||
		f.Fullname == "QSqlTableModel::submit" ||
		f.Fullname == "QFormLayout::itemAt" ||
		f.Fullname == "QGraphicsGridLayout::itemAt" ||
		f.Fullname == "QGridLayout::addItem" ||

		((f.ClassName() == "QGraphicsGridLayout" || f.ClassName() == "QFormLayout") && f.Name == "itemAt" && f.OverloadNumber == "2") {
		return false
	}

	if State.Minimal {
		return f.Export || f.Meta == DESTRUCTOR || f.Fullname == "QObject::destroyed" || strings.HasPrefix(f.Name, TILDE)
	}

	return true
}

func IsBlockedDefault() []string {
	return []string{
		"QAnimationGroup::updateCurrentTime",
		"QAnimationGroup::duration",
		"QAbstractProxyModel::columnCount",
		"QAbstractTableModel::columnCount",
		"QAbstractListModel::data",
		"QAbstractTableModel::data",
		"QAbstractProxyModel::index",
		"QAbstractProxyModel::parent",
		"QAbstractListModel::rowCount",
		"QAbstractProxyModel::rowCount",
		"QAbstractTableModel::rowCount",

		"QNetworkReply::readData",

		"QPagedPaintDevice::paintEngine",
		"QAccessibleObject::childCount",
		"QAccessibleObject::indexOfChild",
		"QAccessibleObject::role",
		"QAccessibleObject::text",
		"QAccessibleObject::child",
		"QAccessibleObject::parent",
		"QAbstractGraphicsShapeItem::paint",
		"QGraphicsObject::paint",
		"QLayout::sizeHint",
		"QAbstractGraphicsShapeItem::boundingRect",
		"QGraphicsObject::boundingRect",
		"QGraphicsLayout::sizeHint",

		"QSimpleXmlNodeModel::typedValue",
		"QSimpleXmlNodeModel::documentUri",
		"QSimpleXmlNodeModel::compareOrder",
		"QSimpleXmlNodeModel::nextFromSimpleAxis",
		"QSimpleXmlNodeModel::kind",
		"QSimpleXmlNodeModel::name",
		"QSimpleXmlNodeModel::root",

		"QAbstractPlanarVideoBuffer::unmap",
		"QAbstractPlanarVideoBuffer::mapMode",

		"QSGDynamicTexture::bind",
		"QSGDynamicTexture::hasMipmaps",
		"QSGDynamicTexture::textureSize",
		"QSGDynamicTexture::hasAlphaChannel",
		"QSGDynamicTexture::textureId",

		"QModbusClient::open",
		"QModbusServer::open",
		"QModbusClient::close",
		"QModbusServer::close",

		"QAbstractBarSeries::type",
		"QXYSeries::type",
	}
}

//TODO: combine
func (f *Function) IsDerivedFromVirtual() bool {
	if f.Virtual != "non" {
		return true
	}

	var class, ok = f.Class()
	if !ok {
		//return false
	}

	for _, bc := range class.GetAllBases() {
		if bclass, ok := State.ClassMap[bc]; ok {

			for _, cf := range bclass.Functions {
				if cf.Name == f.Name &&

					cf.Output == f.Output && len(cf.Parameters) == len(f.Parameters) &&
					cf.Virtual != "non" {

					var similar = true
					for i, cfp := range cf.Parameters {
						if cfp.Value != f.Parameters[i].Value {
							similar = false
						}
					}
					if similar {
						return true
					}
				}
			}

		}
	}

	return false
}

//TODO: combine
func (f *Function) IsDerivedFromImpure() bool {
	if f.Static || f.Virtual == PURE {
		return false
	}

	var class, ok = f.Class()
	if !ok {
		//return false
	}

	if f.Virtual == IMPURE {
		return true
	}

	for _, bc := range class.GetAllBases() {
		if bclass, ok := State.ClassMap[bc]; ok {

			for _, cf := range bclass.Functions {
				if cf.Name == f.Name &&

					cf.Output == f.Output && len(cf.Parameters) == len(f.Parameters) &&
					cf.Virtual == IMPURE {

					var similar = true
					for i, cfp := range cf.Parameters {
						if cfp.Value != f.Parameters[i].Value {
							similar = false
						}
					}
					if similar {
						return true
					}
				}
			}

		}
	}

	return false
}

func (f *Function) IsDerivedFromPure() bool {
	var class, ok = f.Class()
	if !ok {
		//return false
	}

	if f.Virtual == PURE {
		return true
	}

	for _, bc := range class.GetAllBases() {
		if bclass, ok := State.ClassMap[bc]; ok {

			for _, cf := range bclass.Functions {
				if cf.Name == f.Name &&

					cf.Output == f.Output && len(cf.Parameters) == len(f.Parameters) &&
					cf.Virtual == PURE {

					var similar = true
					for i, cfp := range cf.Parameters {
						if cfp.Value != f.Parameters[i].Value {
							similar = false
						}
					}
					if similar {
						return true
					}
				}
			}

		}
	}

	return false
}

func (f *Function) FindDeepestImplementation() string {
	var c, _ = f.Class()

	for _, bcn := range c.GetBases() {
		var bc, ok = State.ClassMap[bcn]
		if !ok {
			continue
		}

		var f = *f
		f.Fullname = fmt.Sprintf("%v::%v", bcn, f.Name)
		var out = f.FindDeepestImplementation()
		if out != "" {
			if c.Module != bc.Module {
				if f.SignalMode == CALLBACK || f.Default || f.Static || strings.HasPrefix(f.Name, "__") {
					return c.Name
				}

				//TODO: --->
				if strings.HasPrefix(f.Name, "__") {
					if f.Root().IsDerivedFromVirtual() {
						return c.Name
					}
				}
				//<--

				f.Fullname = fmt.Sprintf("%v::%v", c.Name, f.Name)
				if plist, _ := f.PossiblePolymorphicDerivations(true); len(plist) > 0 {
					return c.Name
				}
			}
			var lf = bc.GetFunction(f.Name)
			if lf != nil && lf.Virtual == PURE {
				return c.Name
			}
			return out
		}
	}

	if c.HasFunction(f) {
		return c.Name
	}

	return ""
}

func (f *Function) Implements() bool {
	return f.FindDeepestImplementation() == f.ClassName()
}

func (f *Function) Root() *Function {
	var c, ok = f.Class()
	if !ok || !strings.HasPrefix(f.Name, "__") {
		return f
	}

	for _, bcn := range c.GetAllBases() {
		var bc, ok = State.ClassMap[bcn]
		if !ok {
			continue
		}
		for _, cf := range bc.Functions {
			if cf.Name == strings.Split(strings.TrimPrefix(f.Name, "__"), "_")[0] && cf.OverloadNumber == f.OverloadNumber {
				return cf
			}
		}
	}

	return f
}
