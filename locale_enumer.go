// Code generated by "enumer -type=Locale -linecomment -json"; DO NOT EDIT.

package azuretexttospeech

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _LocaleName = "LocaleafZALocaleamETLocalearAELocalearBHLocalearDZLocalearEGLocalearIQLocalearJOLocalearKWLocalearLYLocalearMALocalearQALocalearSALocalearSYLocalearTNLocalearYELocalebgBGLocalebnBDLocalecaESLocalecsCZLocalecyGBLocaledaDKLocaledeATLocaledeCHLocaledeDELocaleelGRLocaleenAULocaleenCALocaleenGBLocaleenHKLocaleenIELocaleenINLocaleenKELocaleenNGLocaleenNZLocaleenPHLocaleenSGLocaleenTZLocaleenUSLocaleenZALocaleesARLocaleesBOLocaleesCLLocaleesCOLocaleesCRLocaleesCULocaleesDOLocaleesECLocaleesESLocaleesGQLocaleesGTLocaleesHNLocaleesMXLocaleesNILocaleesPALocaleesPELocaleesPRLocaleesPYLocaleesSVLocaleesUSLocaleesUYLocaleesVELocaleetEELocalefaIRLocalefiFILocalefilPHLocalefrBELocalefrCALocalefrCHLocalefrFRLocalegaIELocaleglESLocaleguINLocaleheILLocalehiINLocalehrHRLocalehuHULocaleidIDLocaleitITLocalejaJPLocalejvIDLocalekmKHLocalekoKRLocaleltLTLocalelvLVLocalemrINLocalemsMYLocalemtMTLocalemyMMLocalenbNOLocalenlBELocalenlNLLocaleplPLLocaleptBRLocaleptPTLocaleroROLocaleruRULocaleskSKLocaleslSILocalesoSOLocalesuIDLocalesvSELocaleswKELocaleswTZLocaletaINLocaletaLKLocaletaSGLocaleteINLocalethTHLocaletrTRLocaleukUALocaleurINLocaleurPKLocaleuzUZLocaleviVNLocalezhCNLocalezhHKLocalezhTWLocalezuZA"

var _LocaleIndex = [...]uint16{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220, 230, 240, 250, 260, 270, 280, 290, 300, 310, 320, 330, 340, 350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460, 470, 480, 490, 500, 510, 520, 530, 540, 550, 560, 570, 580, 590, 600, 610, 620, 630, 640, 650, 661, 671, 681, 691, 701, 711, 721, 731, 741, 751, 761, 771, 781, 791, 801, 811, 821, 831, 841, 851, 861, 871, 881, 891, 901, 911, 921, 931, 941, 951, 961, 971, 981, 991, 1001, 1011, 1021, 1031, 1041, 1051, 1061, 1071, 1081, 1091, 1101, 1111, 1121, 1131, 1141, 1151, 1161, 1171, 1181, 1191}

const _LocaleLowerName = "localeafzalocaleametlocalearaelocalearbhlocaleardzlocaleareglocaleariqlocalearjolocalearkwlocalearlylocalearmalocalearqalocalearsalocalearsylocaleartnlocalearyelocalebgbglocalebnbdlocalecaeslocalecsczlocalecygblocaledadklocaledeatlocaledechlocalededelocaleelgrlocaleenaulocaleencalocaleengblocaleenhklocaleenielocaleeninlocaleenkelocaleennglocaleennzlocaleenphlocaleensglocaleentzlocaleenuslocaleenzalocaleesarlocaleesbolocaleescllocaleescolocaleescrlocaleesculocaleesdolocaleeseclocaleeseslocaleesgqlocaleesgtlocaleeshnlocaleesmxlocaleesnilocaleespalocaleespelocaleesprlocaleespylocaleessvlocaleesuslocaleesuylocaleesvelocaleeteelocalefairlocalefifilocalefilphlocalefrbelocalefrcalocalefrchlocalefrfrlocalegaielocalegleslocaleguinlocaleheillocalehiinlocalehrhrlocalehuhulocaleididlocaleititlocalejajplocalejvidlocalekmkhlocalekokrlocaleltltlocalelvlvlocalemrinlocalemsmylocalemtmtlocalemymmlocalenbnolocalenlbelocalenlnllocaleplpllocaleptbrlocaleptptlocalerorolocalerurulocalesksklocaleslsilocalesosolocalesuidlocalesvselocaleswkelocaleswtzlocaletainlocaletalklocaletasglocaleteinlocaleththlocaletrtrlocaleukualocaleurinlocaleurpklocaleuzuzlocalevivnlocalezhcnlocalezhhklocalezhtwlocalezuza"

func (i Locale) String() string {
	if i < 0 || i >= Locale(len(_LocaleIndex)-1) {
		return fmt.Sprintf("Locale(%d)", i)
	}
	return _LocaleName[_LocaleIndex[i]:_LocaleIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _LocaleNoOp() {
	var x [1]struct{}
	_ = x[LocaleafZA-(0)]
	_ = x[LocaleamET-(1)]
	_ = x[LocalearAE-(2)]
	_ = x[LocalearBH-(3)]
	_ = x[LocalearDZ-(4)]
	_ = x[LocalearEG-(5)]
	_ = x[LocalearIQ-(6)]
	_ = x[LocalearJO-(7)]
	_ = x[LocalearKW-(8)]
	_ = x[LocalearLY-(9)]
	_ = x[LocalearMA-(10)]
	_ = x[LocalearQA-(11)]
	_ = x[LocalearSA-(12)]
	_ = x[LocalearSY-(13)]
	_ = x[LocalearTN-(14)]
	_ = x[LocalearYE-(15)]
	_ = x[LocalebgBG-(16)]
	_ = x[LocalebnBD-(17)]
	_ = x[LocalecaES-(18)]
	_ = x[LocalecsCZ-(19)]
	_ = x[LocalecyGB-(20)]
	_ = x[LocaledaDK-(21)]
	_ = x[LocaledeAT-(22)]
	_ = x[LocaledeCH-(23)]
	_ = x[LocaledeDE-(24)]
	_ = x[LocaleelGR-(25)]
	_ = x[LocaleenAU-(26)]
	_ = x[LocaleenCA-(27)]
	_ = x[LocaleenGB-(28)]
	_ = x[LocaleenHK-(29)]
	_ = x[LocaleenIE-(30)]
	_ = x[LocaleenIN-(31)]
	_ = x[LocaleenKE-(32)]
	_ = x[LocaleenNG-(33)]
	_ = x[LocaleenNZ-(34)]
	_ = x[LocaleenPH-(35)]
	_ = x[LocaleenSG-(36)]
	_ = x[LocaleenTZ-(37)]
	_ = x[LocaleenUS-(38)]
	_ = x[LocaleenZA-(39)]
	_ = x[LocaleesAR-(40)]
	_ = x[LocaleesBO-(41)]
	_ = x[LocaleesCL-(42)]
	_ = x[LocaleesCO-(43)]
	_ = x[LocaleesCR-(44)]
	_ = x[LocaleesCU-(45)]
	_ = x[LocaleesDO-(46)]
	_ = x[LocaleesEC-(47)]
	_ = x[LocaleesES-(48)]
	_ = x[LocaleesGQ-(49)]
	_ = x[LocaleesGT-(50)]
	_ = x[LocaleesHN-(51)]
	_ = x[LocaleesMX-(52)]
	_ = x[LocaleesNI-(53)]
	_ = x[LocaleesPA-(54)]
	_ = x[LocaleesPE-(55)]
	_ = x[LocaleesPR-(56)]
	_ = x[LocaleesPY-(57)]
	_ = x[LocaleesSV-(58)]
	_ = x[LocaleesUS-(59)]
	_ = x[LocaleesUY-(60)]
	_ = x[LocaleesVE-(61)]
	_ = x[LocaleetEE-(62)]
	_ = x[LocalefaIR-(63)]
	_ = x[LocalefiFI-(64)]
	_ = x[LocalefilPH-(65)]
	_ = x[LocalefrBE-(66)]
	_ = x[LocalefrCA-(67)]
	_ = x[LocalefrCH-(68)]
	_ = x[LocalefrFR-(69)]
	_ = x[LocalegaIE-(70)]
	_ = x[LocaleglES-(71)]
	_ = x[LocaleguIN-(72)]
	_ = x[LocaleheIL-(73)]
	_ = x[LocalehiIN-(74)]
	_ = x[LocalehrHR-(75)]
	_ = x[LocalehuHU-(76)]
	_ = x[LocaleidID-(77)]
	_ = x[LocaleitIT-(78)]
	_ = x[LocalejaJP-(79)]
	_ = x[LocalejvID-(80)]
	_ = x[LocalekmKH-(81)]
	_ = x[LocalekoKR-(82)]
	_ = x[LocaleltLT-(83)]
	_ = x[LocalelvLV-(84)]
	_ = x[LocalemrIN-(85)]
	_ = x[LocalemsMY-(86)]
	_ = x[LocalemtMT-(87)]
	_ = x[LocalemyMM-(88)]
	_ = x[LocalenbNO-(89)]
	_ = x[LocalenlBE-(90)]
	_ = x[LocalenlNL-(91)]
	_ = x[LocaleplPL-(92)]
	_ = x[LocaleptBR-(93)]
	_ = x[LocaleptPT-(94)]
	_ = x[LocaleroRO-(95)]
	_ = x[LocaleruRU-(96)]
	_ = x[LocaleskSK-(97)]
	_ = x[LocaleslSI-(98)]
	_ = x[LocalesoSO-(99)]
	_ = x[LocalesuID-(100)]
	_ = x[LocalesvSE-(101)]
	_ = x[LocaleswKE-(102)]
	_ = x[LocaleswTZ-(103)]
	_ = x[LocaletaIN-(104)]
	_ = x[LocaletaLK-(105)]
	_ = x[LocaletaSG-(106)]
	_ = x[LocaleteIN-(107)]
	_ = x[LocalethTH-(108)]
	_ = x[LocaletrTR-(109)]
	_ = x[LocaleukUA-(110)]
	_ = x[LocaleurIN-(111)]
	_ = x[LocaleurPK-(112)]
	_ = x[LocaleuzUZ-(113)]
	_ = x[LocaleviVN-(114)]
	_ = x[LocalezhCN-(115)]
	_ = x[LocalezhHK-(116)]
	_ = x[LocalezhTW-(117)]
	_ = x[LocalezuZA-(118)]
}

var _LocaleValues = []Locale{LocaleafZA, LocaleamET, LocalearAE, LocalearBH, LocalearDZ, LocalearEG, LocalearIQ, LocalearJO, LocalearKW, LocalearLY, LocalearMA, LocalearQA, LocalearSA, LocalearSY, LocalearTN, LocalearYE, LocalebgBG, LocalebnBD, LocalecaES, LocalecsCZ, LocalecyGB, LocaledaDK, LocaledeAT, LocaledeCH, LocaledeDE, LocaleelGR, LocaleenAU, LocaleenCA, LocaleenGB, LocaleenHK, LocaleenIE, LocaleenIN, LocaleenKE, LocaleenNG, LocaleenNZ, LocaleenPH, LocaleenSG, LocaleenTZ, LocaleenUS, LocaleenZA, LocaleesAR, LocaleesBO, LocaleesCL, LocaleesCO, LocaleesCR, LocaleesCU, LocaleesDO, LocaleesEC, LocaleesES, LocaleesGQ, LocaleesGT, LocaleesHN, LocaleesMX, LocaleesNI, LocaleesPA, LocaleesPE, LocaleesPR, LocaleesPY, LocaleesSV, LocaleesUS, LocaleesUY, LocaleesVE, LocaleetEE, LocalefaIR, LocalefiFI, LocalefilPH, LocalefrBE, LocalefrCA, LocalefrCH, LocalefrFR, LocalegaIE, LocaleglES, LocaleguIN, LocaleheIL, LocalehiIN, LocalehrHR, LocalehuHU, LocaleidID, LocaleitIT, LocalejaJP, LocalejvID, LocalekmKH, LocalekoKR, LocaleltLT, LocalelvLV, LocalemrIN, LocalemsMY, LocalemtMT, LocalemyMM, LocalenbNO, LocalenlBE, LocalenlNL, LocaleplPL, LocaleptBR, LocaleptPT, LocaleroRO, LocaleruRU, LocaleskSK, LocaleslSI, LocalesoSO, LocalesuID, LocalesvSE, LocaleswKE, LocaleswTZ, LocaletaIN, LocaletaLK, LocaletaSG, LocaleteIN, LocalethTH, LocaletrTR, LocaleukUA, LocaleurIN, LocaleurPK, LocaleuzUZ, LocaleviVN, LocalezhCN, LocalezhHK, LocalezhTW, LocalezuZA}

var _LocaleNameToValueMap = map[string]Locale{
	_LocaleName[0:10]:           LocaleafZA,
	_LocaleLowerName[0:10]:      LocaleafZA,
	_LocaleName[10:20]:          LocaleamET,
	_LocaleLowerName[10:20]:     LocaleamET,
	_LocaleName[20:30]:          LocalearAE,
	_LocaleLowerName[20:30]:     LocalearAE,
	_LocaleName[30:40]:          LocalearBH,
	_LocaleLowerName[30:40]:     LocalearBH,
	_LocaleName[40:50]:          LocalearDZ,
	_LocaleLowerName[40:50]:     LocalearDZ,
	_LocaleName[50:60]:          LocalearEG,
	_LocaleLowerName[50:60]:     LocalearEG,
	_LocaleName[60:70]:          LocalearIQ,
	_LocaleLowerName[60:70]:     LocalearIQ,
	_LocaleName[70:80]:          LocalearJO,
	_LocaleLowerName[70:80]:     LocalearJO,
	_LocaleName[80:90]:          LocalearKW,
	_LocaleLowerName[80:90]:     LocalearKW,
	_LocaleName[90:100]:         LocalearLY,
	_LocaleLowerName[90:100]:    LocalearLY,
	_LocaleName[100:110]:        LocalearMA,
	_LocaleLowerName[100:110]:   LocalearMA,
	_LocaleName[110:120]:        LocalearQA,
	_LocaleLowerName[110:120]:   LocalearQA,
	_LocaleName[120:130]:        LocalearSA,
	_LocaleLowerName[120:130]:   LocalearSA,
	_LocaleName[130:140]:        LocalearSY,
	_LocaleLowerName[130:140]:   LocalearSY,
	_LocaleName[140:150]:        LocalearTN,
	_LocaleLowerName[140:150]:   LocalearTN,
	_LocaleName[150:160]:        LocalearYE,
	_LocaleLowerName[150:160]:   LocalearYE,
	_LocaleName[160:170]:        LocalebgBG,
	_LocaleLowerName[160:170]:   LocalebgBG,
	_LocaleName[170:180]:        LocalebnBD,
	_LocaleLowerName[170:180]:   LocalebnBD,
	_LocaleName[180:190]:        LocalecaES,
	_LocaleLowerName[180:190]:   LocalecaES,
	_LocaleName[190:200]:        LocalecsCZ,
	_LocaleLowerName[190:200]:   LocalecsCZ,
	_LocaleName[200:210]:        LocalecyGB,
	_LocaleLowerName[200:210]:   LocalecyGB,
	_LocaleName[210:220]:        LocaledaDK,
	_LocaleLowerName[210:220]:   LocaledaDK,
	_LocaleName[220:230]:        LocaledeAT,
	_LocaleLowerName[220:230]:   LocaledeAT,
	_LocaleName[230:240]:        LocaledeCH,
	_LocaleLowerName[230:240]:   LocaledeCH,
	_LocaleName[240:250]:        LocaledeDE,
	_LocaleLowerName[240:250]:   LocaledeDE,
	_LocaleName[250:260]:        LocaleelGR,
	_LocaleLowerName[250:260]:   LocaleelGR,
	_LocaleName[260:270]:        LocaleenAU,
	_LocaleLowerName[260:270]:   LocaleenAU,
	_LocaleName[270:280]:        LocaleenCA,
	_LocaleLowerName[270:280]:   LocaleenCA,
	_LocaleName[280:290]:        LocaleenGB,
	_LocaleLowerName[280:290]:   LocaleenGB,
	_LocaleName[290:300]:        LocaleenHK,
	_LocaleLowerName[290:300]:   LocaleenHK,
	_LocaleName[300:310]:        LocaleenIE,
	_LocaleLowerName[300:310]:   LocaleenIE,
	_LocaleName[310:320]:        LocaleenIN,
	_LocaleLowerName[310:320]:   LocaleenIN,
	_LocaleName[320:330]:        LocaleenKE,
	_LocaleLowerName[320:330]:   LocaleenKE,
	_LocaleName[330:340]:        LocaleenNG,
	_LocaleLowerName[330:340]:   LocaleenNG,
	_LocaleName[340:350]:        LocaleenNZ,
	_LocaleLowerName[340:350]:   LocaleenNZ,
	_LocaleName[350:360]:        LocaleenPH,
	_LocaleLowerName[350:360]:   LocaleenPH,
	_LocaleName[360:370]:        LocaleenSG,
	_LocaleLowerName[360:370]:   LocaleenSG,
	_LocaleName[370:380]:        LocaleenTZ,
	_LocaleLowerName[370:380]:   LocaleenTZ,
	_LocaleName[380:390]:        LocaleenUS,
	_LocaleLowerName[380:390]:   LocaleenUS,
	_LocaleName[390:400]:        LocaleenZA,
	_LocaleLowerName[390:400]:   LocaleenZA,
	_LocaleName[400:410]:        LocaleesAR,
	_LocaleLowerName[400:410]:   LocaleesAR,
	_LocaleName[410:420]:        LocaleesBO,
	_LocaleLowerName[410:420]:   LocaleesBO,
	_LocaleName[420:430]:        LocaleesCL,
	_LocaleLowerName[420:430]:   LocaleesCL,
	_LocaleName[430:440]:        LocaleesCO,
	_LocaleLowerName[430:440]:   LocaleesCO,
	_LocaleName[440:450]:        LocaleesCR,
	_LocaleLowerName[440:450]:   LocaleesCR,
	_LocaleName[450:460]:        LocaleesCU,
	_LocaleLowerName[450:460]:   LocaleesCU,
	_LocaleName[460:470]:        LocaleesDO,
	_LocaleLowerName[460:470]:   LocaleesDO,
	_LocaleName[470:480]:        LocaleesEC,
	_LocaleLowerName[470:480]:   LocaleesEC,
	_LocaleName[480:490]:        LocaleesES,
	_LocaleLowerName[480:490]:   LocaleesES,
	_LocaleName[490:500]:        LocaleesGQ,
	_LocaleLowerName[490:500]:   LocaleesGQ,
	_LocaleName[500:510]:        LocaleesGT,
	_LocaleLowerName[500:510]:   LocaleesGT,
	_LocaleName[510:520]:        LocaleesHN,
	_LocaleLowerName[510:520]:   LocaleesHN,
	_LocaleName[520:530]:        LocaleesMX,
	_LocaleLowerName[520:530]:   LocaleesMX,
	_LocaleName[530:540]:        LocaleesNI,
	_LocaleLowerName[530:540]:   LocaleesNI,
	_LocaleName[540:550]:        LocaleesPA,
	_LocaleLowerName[540:550]:   LocaleesPA,
	_LocaleName[550:560]:        LocaleesPE,
	_LocaleLowerName[550:560]:   LocaleesPE,
	_LocaleName[560:570]:        LocaleesPR,
	_LocaleLowerName[560:570]:   LocaleesPR,
	_LocaleName[570:580]:        LocaleesPY,
	_LocaleLowerName[570:580]:   LocaleesPY,
	_LocaleName[580:590]:        LocaleesSV,
	_LocaleLowerName[580:590]:   LocaleesSV,
	_LocaleName[590:600]:        LocaleesUS,
	_LocaleLowerName[590:600]:   LocaleesUS,
	_LocaleName[600:610]:        LocaleesUY,
	_LocaleLowerName[600:610]:   LocaleesUY,
	_LocaleName[610:620]:        LocaleesVE,
	_LocaleLowerName[610:620]:   LocaleesVE,
	_LocaleName[620:630]:        LocaleetEE,
	_LocaleLowerName[620:630]:   LocaleetEE,
	_LocaleName[630:640]:        LocalefaIR,
	_LocaleLowerName[630:640]:   LocalefaIR,
	_LocaleName[640:650]:        LocalefiFI,
	_LocaleLowerName[640:650]:   LocalefiFI,
	_LocaleName[650:661]:        LocalefilPH,
	_LocaleLowerName[650:661]:   LocalefilPH,
	_LocaleName[661:671]:        LocalefrBE,
	_LocaleLowerName[661:671]:   LocalefrBE,
	_LocaleName[671:681]:        LocalefrCA,
	_LocaleLowerName[671:681]:   LocalefrCA,
	_LocaleName[681:691]:        LocalefrCH,
	_LocaleLowerName[681:691]:   LocalefrCH,
	_LocaleName[691:701]:        LocalefrFR,
	_LocaleLowerName[691:701]:   LocalefrFR,
	_LocaleName[701:711]:        LocalegaIE,
	_LocaleLowerName[701:711]:   LocalegaIE,
	_LocaleName[711:721]:        LocaleglES,
	_LocaleLowerName[711:721]:   LocaleglES,
	_LocaleName[721:731]:        LocaleguIN,
	_LocaleLowerName[721:731]:   LocaleguIN,
	_LocaleName[731:741]:        LocaleheIL,
	_LocaleLowerName[731:741]:   LocaleheIL,
	_LocaleName[741:751]:        LocalehiIN,
	_LocaleLowerName[741:751]:   LocalehiIN,
	_LocaleName[751:761]:        LocalehrHR,
	_LocaleLowerName[751:761]:   LocalehrHR,
	_LocaleName[761:771]:        LocalehuHU,
	_LocaleLowerName[761:771]:   LocalehuHU,
	_LocaleName[771:781]:        LocaleidID,
	_LocaleLowerName[771:781]:   LocaleidID,
	_LocaleName[781:791]:        LocaleitIT,
	_LocaleLowerName[781:791]:   LocaleitIT,
	_LocaleName[791:801]:        LocalejaJP,
	_LocaleLowerName[791:801]:   LocalejaJP,
	_LocaleName[801:811]:        LocalejvID,
	_LocaleLowerName[801:811]:   LocalejvID,
	_LocaleName[811:821]:        LocalekmKH,
	_LocaleLowerName[811:821]:   LocalekmKH,
	_LocaleName[821:831]:        LocalekoKR,
	_LocaleLowerName[821:831]:   LocalekoKR,
	_LocaleName[831:841]:        LocaleltLT,
	_LocaleLowerName[831:841]:   LocaleltLT,
	_LocaleName[841:851]:        LocalelvLV,
	_LocaleLowerName[841:851]:   LocalelvLV,
	_LocaleName[851:861]:        LocalemrIN,
	_LocaleLowerName[851:861]:   LocalemrIN,
	_LocaleName[861:871]:        LocalemsMY,
	_LocaleLowerName[861:871]:   LocalemsMY,
	_LocaleName[871:881]:        LocalemtMT,
	_LocaleLowerName[871:881]:   LocalemtMT,
	_LocaleName[881:891]:        LocalemyMM,
	_LocaleLowerName[881:891]:   LocalemyMM,
	_LocaleName[891:901]:        LocalenbNO,
	_LocaleLowerName[891:901]:   LocalenbNO,
	_LocaleName[901:911]:        LocalenlBE,
	_LocaleLowerName[901:911]:   LocalenlBE,
	_LocaleName[911:921]:        LocalenlNL,
	_LocaleLowerName[911:921]:   LocalenlNL,
	_LocaleName[921:931]:        LocaleplPL,
	_LocaleLowerName[921:931]:   LocaleplPL,
	_LocaleName[931:941]:        LocaleptBR,
	_LocaleLowerName[931:941]:   LocaleptBR,
	_LocaleName[941:951]:        LocaleptPT,
	_LocaleLowerName[941:951]:   LocaleptPT,
	_LocaleName[951:961]:        LocaleroRO,
	_LocaleLowerName[951:961]:   LocaleroRO,
	_LocaleName[961:971]:        LocaleruRU,
	_LocaleLowerName[961:971]:   LocaleruRU,
	_LocaleName[971:981]:        LocaleskSK,
	_LocaleLowerName[971:981]:   LocaleskSK,
	_LocaleName[981:991]:        LocaleslSI,
	_LocaleLowerName[981:991]:   LocaleslSI,
	_LocaleName[991:1001]:       LocalesoSO,
	_LocaleLowerName[991:1001]:  LocalesoSO,
	_LocaleName[1001:1011]:      LocalesuID,
	_LocaleLowerName[1001:1011]: LocalesuID,
	_LocaleName[1011:1021]:      LocalesvSE,
	_LocaleLowerName[1011:1021]: LocalesvSE,
	_LocaleName[1021:1031]:      LocaleswKE,
	_LocaleLowerName[1021:1031]: LocaleswKE,
	_LocaleName[1031:1041]:      LocaleswTZ,
	_LocaleLowerName[1031:1041]: LocaleswTZ,
	_LocaleName[1041:1051]:      LocaletaIN,
	_LocaleLowerName[1041:1051]: LocaletaIN,
	_LocaleName[1051:1061]:      LocaletaLK,
	_LocaleLowerName[1051:1061]: LocaletaLK,
	_LocaleName[1061:1071]:      LocaletaSG,
	_LocaleLowerName[1061:1071]: LocaletaSG,
	_LocaleName[1071:1081]:      LocaleteIN,
	_LocaleLowerName[1071:1081]: LocaleteIN,
	_LocaleName[1081:1091]:      LocalethTH,
	_LocaleLowerName[1081:1091]: LocalethTH,
	_LocaleName[1091:1101]:      LocaletrTR,
	_LocaleLowerName[1091:1101]: LocaletrTR,
	_LocaleName[1101:1111]:      LocaleukUA,
	_LocaleLowerName[1101:1111]: LocaleukUA,
	_LocaleName[1111:1121]:      LocaleurIN,
	_LocaleLowerName[1111:1121]: LocaleurIN,
	_LocaleName[1121:1131]:      LocaleurPK,
	_LocaleLowerName[1121:1131]: LocaleurPK,
	_LocaleName[1131:1141]:      LocaleuzUZ,
	_LocaleLowerName[1131:1141]: LocaleuzUZ,
	_LocaleName[1141:1151]:      LocaleviVN,
	_LocaleLowerName[1141:1151]: LocaleviVN,
	_LocaleName[1151:1161]:      LocalezhCN,
	_LocaleLowerName[1151:1161]: LocalezhCN,
	_LocaleName[1161:1171]:      LocalezhHK,
	_LocaleLowerName[1161:1171]: LocalezhHK,
	_LocaleName[1171:1181]:      LocalezhTW,
	_LocaleLowerName[1171:1181]: LocalezhTW,
	_LocaleName[1181:1191]:      LocalezuZA,
	_LocaleLowerName[1181:1191]: LocalezuZA,
}

var _LocaleNames = []string{
	_LocaleName[0:10],
	_LocaleName[10:20],
	_LocaleName[20:30],
	_LocaleName[30:40],
	_LocaleName[40:50],
	_LocaleName[50:60],
	_LocaleName[60:70],
	_LocaleName[70:80],
	_LocaleName[80:90],
	_LocaleName[90:100],
	_LocaleName[100:110],
	_LocaleName[110:120],
	_LocaleName[120:130],
	_LocaleName[130:140],
	_LocaleName[140:150],
	_LocaleName[150:160],
	_LocaleName[160:170],
	_LocaleName[170:180],
	_LocaleName[180:190],
	_LocaleName[190:200],
	_LocaleName[200:210],
	_LocaleName[210:220],
	_LocaleName[220:230],
	_LocaleName[230:240],
	_LocaleName[240:250],
	_LocaleName[250:260],
	_LocaleName[260:270],
	_LocaleName[270:280],
	_LocaleName[280:290],
	_LocaleName[290:300],
	_LocaleName[300:310],
	_LocaleName[310:320],
	_LocaleName[320:330],
	_LocaleName[330:340],
	_LocaleName[340:350],
	_LocaleName[350:360],
	_LocaleName[360:370],
	_LocaleName[370:380],
	_LocaleName[380:390],
	_LocaleName[390:400],
	_LocaleName[400:410],
	_LocaleName[410:420],
	_LocaleName[420:430],
	_LocaleName[430:440],
	_LocaleName[440:450],
	_LocaleName[450:460],
	_LocaleName[460:470],
	_LocaleName[470:480],
	_LocaleName[480:490],
	_LocaleName[490:500],
	_LocaleName[500:510],
	_LocaleName[510:520],
	_LocaleName[520:530],
	_LocaleName[530:540],
	_LocaleName[540:550],
	_LocaleName[550:560],
	_LocaleName[560:570],
	_LocaleName[570:580],
	_LocaleName[580:590],
	_LocaleName[590:600],
	_LocaleName[600:610],
	_LocaleName[610:620],
	_LocaleName[620:630],
	_LocaleName[630:640],
	_LocaleName[640:650],
	_LocaleName[650:661],
	_LocaleName[661:671],
	_LocaleName[671:681],
	_LocaleName[681:691],
	_LocaleName[691:701],
	_LocaleName[701:711],
	_LocaleName[711:721],
	_LocaleName[721:731],
	_LocaleName[731:741],
	_LocaleName[741:751],
	_LocaleName[751:761],
	_LocaleName[761:771],
	_LocaleName[771:781],
	_LocaleName[781:791],
	_LocaleName[791:801],
	_LocaleName[801:811],
	_LocaleName[811:821],
	_LocaleName[821:831],
	_LocaleName[831:841],
	_LocaleName[841:851],
	_LocaleName[851:861],
	_LocaleName[861:871],
	_LocaleName[871:881],
	_LocaleName[881:891],
	_LocaleName[891:901],
	_LocaleName[901:911],
	_LocaleName[911:921],
	_LocaleName[921:931],
	_LocaleName[931:941],
	_LocaleName[941:951],
	_LocaleName[951:961],
	_LocaleName[961:971],
	_LocaleName[971:981],
	_LocaleName[981:991],
	_LocaleName[991:1001],
	_LocaleName[1001:1011],
	_LocaleName[1011:1021],
	_LocaleName[1021:1031],
	_LocaleName[1031:1041],
	_LocaleName[1041:1051],
	_LocaleName[1051:1061],
	_LocaleName[1061:1071],
	_LocaleName[1071:1081],
	_LocaleName[1081:1091],
	_LocaleName[1091:1101],
	_LocaleName[1101:1111],
	_LocaleName[1111:1121],
	_LocaleName[1121:1131],
	_LocaleName[1131:1141],
	_LocaleName[1141:1151],
	_LocaleName[1151:1161],
	_LocaleName[1161:1171],
	_LocaleName[1171:1181],
	_LocaleName[1181:1191],
}

// LocaleString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func LocaleString(s string) (Locale, error) {
	if val, ok := _LocaleNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _LocaleNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Locale values", s)
}

// LocaleValues returns all values of the enum
func LocaleValues() []Locale {
	return _LocaleValues
}

// LocaleStrings returns a slice of all String values of the enum
func LocaleStrings() []string {
	strs := make([]string, len(_LocaleNames))
	copy(strs, _LocaleNames)
	return strs
}

// IsALocale returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Locale) IsALocale() bool {
	for _, v := range _LocaleValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Locale
func (i Locale) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Locale
func (i *Locale) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Locale should be a string, got %s", data)
	}

	var err error
	*i, err = LocaleString(s)
	return err
}
