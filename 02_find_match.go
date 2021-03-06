package main

import (
	"fmt"
	"strings"
)

func part1(ss []string) int {
	n2, n3 := 0, 0
	for _, s := range ss {
		fs := map[rune]int{}
		for _, l := range s {
			fs[l]++
		}
		fsn := map[int]int{}
		for _, v := range fs {
			fsn[v]++
		}
		if fsn[2] > 0 {
			n2++
		}
		if fsn[3] > 0 {
			n3++
		}
	}
	return n2 * n3
}

/*

// O(stringlen * n^2)

func part2(ss []string) string {
	for i, s1 := range ss {
		for _, s2 := range ss[i+1:] {
			diff := 0
			rs1 := []rune(s1)
			rs2 := []rune(s2)
			for k, r := range rs1 {
				if r != rs2[k] {
					diff++
				}
			}
			if diff == 1 {
				out := make([]rune, 0, len(rs1)-1)
				for k, r := range rs1 {
					if r == rs2[k] {
						out = append(out, r)
					}
				}
				return string(out)
			}
		}
	}
	return "???"
}

/*/

// O(stringlen * n), inspired by https://github.com/petertseng

func part2(ss []string) string {
	seen := map[[2]string]bool{}
	for _, s := range ss {
		rs := []rune(s)
		for i := 0; i < len(s); i++ {
			p := [2]string{string(rs[:i]), string(rs[i+1:])}
			if seen[p] {
				return p[0] + p[1]
			}
			seen[p] = true
		}
	}
	return "???"
}

//*/

func main() {
	for i, in := range ins {
		fmt.Println("=== for", i, "===")
		ss := strings.Fields(in)
		fmt.Println("part 1:", part1(ss))
		fmt.Println("part 2:", part2(ss))
		fmt.Println()
	}
}

var ins = map[string]string{
	"github": `umdryebvlapkozostecnihjexg
amdryebalapkozfstwcnrhjqxg
umdcyebvlapaozfstwcnihjqgg
ymdryrbvlapkozfstwcuihjqxg
umdrsebvlapkozxstwcnihjqig
umdryibvlapkohfstwcnfhjqxg
umdryebvqapkozfatwcnihjqxs
umzrpebvlapkozfshwcnihjqxg
fmhryebvlapkozfstwckihjqxg
umdryebvlahkozfstwcnizjrxg
qmdryebvlapkozfslwcnihgqxg
umdiyebjlapknzfstwcnihjqxg
umdryebvlapkoqfstwcaihvqxg
cmdryebvlapkpzfstwcnihjvxg
umdryebvlakkozfstwcgihjixg
umdryebvlasjozfstwcnihqqxg
umdryebvladkozfsvwcnifjqxg
umdrlebvlapaozfstwcniwjqxg
umdryebvlhpkozrstwsnihjqxg
umdryebvcapkozfqtwcnihjrxg
ubdrykbvlapkowfstwcnihjqxg
umdryebvldpkozfstwcnihtqsg
umdryebvlapaozyutwcnihjqxg
umdryibvlapkozfstdfnihjqxg
umdryebvlapgozkstwznihjqxg
umdrxebvlapkozfstwcngxjqxg
umdryekvlapkozfstwclchjqxg
nmdryebvlapkozjsewcnihjqxg
umdryebvyapkozfstfcniheqxg
umdfyebvlapkozfstwcnhhjpxg
umdryelylupkozfstwcnihjqxg
smdryebvlqpkozfstwcnihjdxg
umdryebvlapaozfsuwcnihjqxc
umdryebvlrzkozrstwcnihjqxg
umdbycbvlapkojfstwcnihjqxg
umdryebvlapkonfstwpnirjqxg
uecryebvlapkozfstwcnihpqxg
uqdryebvltpkozfstwcnihrqxg
umdryebvlqsknzfstwcnihjqxg
cmdryebvlapkocfstwcvihjqxg
umdrkebvlapkozqsfwcnihjqxg
umdryabveapkoifstwcnihjqxg
ummrnehvlapkozfstwcnihjqxg
umdryebvlxpkozfstwqnihjtxg
umdryebvlagkozastwcnihjqxh
umdryebvlatkozzhtwcnihjqxg
umdryebvlcpkozfstwrnihjqvg
umdryebvlapkozfsnwcnrhjcxg
umdzyebvlypkozfstwcnibjqxg
nmdryebvlvpkozbstwcnihjqxg
uwdryebvlipkozfstwcnihvqxg
umdraebvlavkozfstwcnihjqwg
umdeyebvlspbozfstwcnihjqxg
umdryxlvlapkozfstwcnihjqxu
umdryegvlapkqzfstwcnirjqxg
umdrupbvlapkozfstwcnihjqog
imxryebvlapkxzfstwcnihjqxg
umdrfebvlapkozowtwcnihjqxg
umdreebvlapkozmstwczihjqxg
undryebdlapkozbstwcnihjqxg
umdryebvlapkpzfetwcnihjqxb
ymdnyebvlapkozfstwinihjqxg
umdryebvaapkozfstwcnihyqqg
umdryebvlapkzzwsrwcnihjqxg
umdrkebvlapkmzfskwcnihjqxg
umdrmebvlapkozfsvwcnidjqxg
umdlyehvlapkozfstwcnihjqkg
umnryebvlrpkozfstwjnihjqxg
uqdryebvlapxozfsawcnihjqxg
vmdruebvlapkozfstwcnihjqqg
umdryabviapkozistwcnihjqxg
umdryebvlapkzzfstwfnihkqxg
uvdryebvlapkozfsxwcuihjqxg
umdlhebvlapkozfstwcnvhjqxg
umdreebvlapkopfstjcnihjqxg
umdryebvlazkomfstwynihjqxg
kmdryebulapkoznstwcnihjqxg
umdryebvxakkozfstwinihjqxg
ukdryobvlapkozistwcnihjqxg
umdryebveapkozfstwcnthjqgg
mmdrtebvlapcozfstwcnihjqxg
umdryebvlapkolistwnnihjqxg
umdryebxlapkozfatwcnihjqxx
uxdryebvlapkozfstwhniheqxg
ufdryebvzapkozfstwcnbhjqxg
amdryhbvlapkozfstwcnifjqxg
umqryebvlaphozfstwcnihjqxn
umdryebvlapkosfstfcnihjqxe
gmkryebvlapkozfstwcnihjmxg
umdrnebvlkpkozfstwcnihjnxg
umdryebvrapkozfstmcndhjqxg
umdryebvmapkozfstichihjqxg
umdryesvnapkozestwcnihjqxg
umeryhbvlapkozfstfcnihjqxg
umdryedvbapkozfstwcnihqqxg
umdryebllapzozfstwcnihjvxg
umdcyebvlzdkozfstwcnihjqxg
umdrybbvlapkbvfstwcnihjqxg
umdrytbglapkozfsthcnihjqxg
umdryebvlkpkozfsteclihjqxg
umdntebvlapkmzfstwcnihjqxg
lkdryebveapkozfstwcnihjqxg
ymdryubvlapkozfstwbnihjqxg
tmrryebvlapkozfstwcnqhjqxg
umdryeovlaekonfstwcnihjqxg
umiryeuvlapkozfstwcnihjwxg
umdryebvlspvozwstwcnihjqxg
umdrtebvlapkoznxtwcnihjqxg
umvryebvlaphozfstwcnahjqxg
umdryebvlapkozfstinniajqxg
umdryebqlapkozfctwcnihjqxx
umdryebvlapkbzfptwcnihjqvg
umdryabviapkozistwcnihjqxd
umdryrbvlapkezfstscnihjqxg
umhryebvlapkozfstacnihxqxg
umdxyelvlapkozfitwcnihjqxg
umdryevvuapkozfstwcnihtqxg
uydrypbvxapkozfstwcnihjqxg
umdryebvlapkopfstwcnihzqxo
uedryebvlapkozistwceihjqxg
umdiyebvlapkozfgtwcnihjqxv
ymdryebvlapkozfsticniqjqxg
umbrkebvlapkozfslwcnihjqxg
umdryebliapkozbstwcnihjqxg
umvryebolapkozfstwcnihjqig
umdryeavbackozfstwcnihjqxg
umdryfbvlapsozfstwcnihaqxg
umdqyebvlapkozfjtgcnihjqxg
umdrjebvlaqkozfstwcyihjqxg
umdryebklaqkozrstwcnihjqxg
umdryebvpapkozfstwcpihjqjg
uydryebhlawkozfstwcnihjqxg
umdyyebvlapkozfstwcykhjqxg
umdryebvlapkozfstwcnitjnxh
umdzyebvlapkozfstwcnehyqxg
mmcryebvlapkozfstwinihjqxg
umdryebvlapuozfstwmvihjqxg
umdryfbvlapkozqstwcnihjmxg
umdryebslapsozfhtwcnihjqxg
umdtyemvlapmozfstwcnihjqxg
umdrxevvlapkozfytwcnihjqxg
umdahebvlapjozfstwcnihjqxg
umdryebvlapkozfstacnivjqxb
umdryebvlzpkozfjtwcnihjyxg
umdryebvlaqkozfstwcnisjqxu
umdrydbvlapkozfsuwcnihjlxg
umdryebvlapkomrstwcnihjqkg
umdryebvlapcozfstmcnwhjqxg
umdryebvlahkozfstwcibhjqxg
gmdrzebvlapkozlstwcnihjqxg
umdryebvlapkezfsswcnrhjqxg
umdryebvlapkoqfitwcgihjqxg
umdrnebvlapkozfsiwcninjqxg
umdryebvlapkozfsrwckohjqxg
umdryebtlapkomfstwcnihjexg
umdryxbvlapjozfstwcnihoqxg
umdpyebvlapkosustwcnihjqxg
umdryebvlapkvzfawwcnihjqxg
umhnyebvlaikozfstwcnihjqxg
umdryebvlagkozfstvknihjqxg
uodryebjlapkoxfstwcnihjqxg
umdryefdlapkozfstwcnyhjqxg
umprmebvtapkozfstwcnihjqxg
umdhyebvlapoozfstwcnihjqgg
uddryebvidpkozfstwcnihjqxg
umdryebtlapkozfetwfnihjqxg
umdbyebolapkozfstwcoihjqxg
umdryebvlapkonfstwcnihjpxo
umdryebvlapkohfstwcnihjqwk
umdryebolalkkzfstwcnihjqxg
updryebvxapkozfstwcnshjqxg
umdryebvlapkovfktwcnuhjqxg
umdrqrbvlppkozfstwcnihjqxg
umdrylgvlapkozfstwrnihjqxg
umdryebvlapkozfstxcnihbqig
uvdryeevlappozfstwcnihjqxg
zmdryebvlapkozfstwcnihqqxt
umdryebvlapvozfstwenihiqxg
umdryebvlbpkozfsgwcnihjlxg
umdryhbvlapkozfstwcnihtqxw
umdreecvlapkozwstwcnihjqxg
umwryebvlapkoztsmwcnihjqxg
ukdryebvfapkozrstwcnihjqxg
umdrylbdlamkozfstwcnihjqxg
umdryebvlapoozwsmwcnihjqxg
umdryebvlapkozfqtwcnnzjqxg
umdryekvlapktzfstwcnohjqxg
umdryebvlapkozfstwcnihjwqo
umdrrebflapkogfstwcnihjqxg
umdryevvlapkozfztwctihjqxg
umdrybbvlapkozfstwcnihxaxg
umdryebvlapkozfsowcnphjqag
smdryebvlapbozfitwcnihjqxg
umdryebvtapiozfstwcnihjqxe
umdryebjlakkozfstwccihjqxg
umdryebvlapdozfshwckihjqxg
umnryebvlapiozfstwcnihlqxg
umdrycbvlapkjzfsnwcnihjqxg
umdryebvyaprozjstwcnihjqxg
ucdryebvlapkozfstwomihjqxg
umdryebvlagklzfstwcnihjqyg
umdryebvladkozfstwcnihjqjh
umdrwebvlapkozfstwdnicjqxg
umdryebvlapkmzfstwcniheqxr
umdryebvlapkjzfstwcviheqxg
umdrvebvlapkozfstwcbihjqmg
umdrfebvlapkoffstwcnihsqxg
umdryebvtarkazfstwcnihjqxg
umdryebvlapkozfstwcfihjcng
umdryebvlapkktostwcnihjqxg
uedryeevlapkozfstwcniijqxg
bmdryebylapkozfstwcnihjqog
umdryebvlmpkoztstwcnihjqeg
umdryepvlarkohfstwcnihjqxg
uwdryebvlapklzfstzcnihjqxg
umdryebklapkozfsswcbihjqxg
umdtyeavlapkozfstwsnihjqxg
umdryebvaapkozfhtfcnihjqxg
umdrpebvlapuozfstwvnihjqxg
umdryebvlapkozffmwcniijqxg
uqdpyebvlapkozfstwfnihjqxg
umdryebvlapuozdstwcnihjhxg
tmdryhbvlapkozfptwcnihjqxg
umdryevvmapkozfstwcnihjgxg
umdryeuvlapmozfstwcnihjwxg
umdryebqlzpkozfbtwcnihjqxg
umdryebvsapkozystwcniqjqxg
imdryebvlapkozfscwinihjqxg
umdryebvlzpkozustwcnmhjqxg
umdrypbvlapbozfsnwcnihjqxg
bmdryebvlapqoznstwcnihjqxg
umdrfebvlapaozfstwcnihxqxg
umdiyebvxapkozfstwcnchjqxg
umdrygbvlapkozfstwcnizjqxz
amdryedvlapkozfstwcnihfqxg
umdryebvvapzozfstwcnihjgxg
undryebvlapkzzfstjcnihjqxg
umdryvbvlapgozfrtwcnihjqxg
umdrkebvlapkozfstwcnihihxg
umdryebvrppkozfsowcnihjqxg
umdryebvlapktzfsdwclihjqxg
otdrdebvlapkozfstwcnihjqxg
mmdryebvlazkozfxtwcnihjqxg
umdryebvlapkozfsbwtnihjqxa
imqryebvrapkozfstwcnihjqxg
umdryebvlrpkozfscwcnihjqlg
uedryebvlapkoznsvwcnihjqxg
umdryebvlqpkozfstscnihjqxj
umerycbvlapkozfstwcnihjqxh
umdkykbvlapjozfstwcnihjqxg`,
	"google": `icxjubroqtunzeyzpomfksahgw
ibxjvbroqtunleyzjdmfksahow
icxjvbroqtinleyzpdmflsahnw
icxjvbnoqtunleyvpgmfksahgw
wcxjvbroqrunleyzpdmfksahge
icxjtbroqtjzleyzpdmfksahgw
icxjvbrohtunleyzpdmfkbahsw
xcxjvbroqcunleyzpdmfksaxgw
ycxjvbroqtunleyzpowfksahgw
icxfvbroqtunleyzpdmfksncgw
ixxjvbuoqtunleyzpdvfksahgw
icfjvbroqtunleyzpdmfksadgt
icxjvbroqdunleyzpdafksqhgw
icxjvbrobtunlelzpdmfkuahgs
ujxjvbroqtunleyzpdmqksahgw
icqjvsroqtunleyzpdmfksahuw
icxjvbroptpnleyzpdmfksangw
ipxjvbroqtunleyzpdmfesahgi
icajvbroqtunltyzpdqfksahgw
ickjvbroqtuzleyzpdmfgsahgw
icxjvbroqtunledzpdmwksahgz
icxjvlroqtsnleyzpdmfksvhgw
icxjvbroqtunleyzpdsfysahvw
icxjvbroqtunwnyzydmfksahgw
ionjvbroqtunleyzpdmfksahgj
icxjvwriqtunleyzpdmfksahgi
ocxjvbroztunleyzpdmfksapgw
icxjvbroqtmnlewzpumfksahgw
ucxjvbroqtunleyzpdmzktahgw
icxgvbroqtunleyztdmfktahgw
icxhvbroqttnleybpdmfksahgw
icxjvbroqtugleyzpdxfkeahgw
acxjvbroqvunlexzpdmfksahgw
icxjvbroqglnleyzpbmfksahgw
icxjvbriqtnvleyzpdmfksahgw
icxjvbreqtunlakzpdmfksahgw
gcxjvbuoqtunleyzpdmfksawgw
icxjvbroqtunleyzpddfkyzhgw
icxjvbjoqtunleyzpdmfqsahhw
icxjvjroqtunleyzpnmfksajgw
ycxjvbroqtunmeyzcdmfksahgw
irxkvbryqtunleyzpdmfksahgw
isxjvbrlqtunleyzpdmsksahgw
icxjvbcoqtunleyzpdfkksahgw
ixnjvbroqtunleyzpdmfkqahgw
wcxjvbroqhunleyzqdmfksahgw
icljvurmqtunleyzpdmfksahgw
ibxjvbroqtayleyzpdmfksahgw
arxjvbroqiunleyzpdmfksahgw
iuxjvbroqtunluyzpdmoksahgw
icxjvbrmqtunleyzpdmfosahew
isxjvbroqtunleyrpdmfksrhgw
icxjvxrpqtunleyzpdmfkdahgw
ichjvbrogtunllyzpdmfksahgw
icxjvbeoqtunlryzpdmfksakgw
icxjvbroqtcnemyzpdmfksahgw
icxjvbroqtybledzpdmfksahgw
icxjvbrqqtunleyzpdmfksgngw
icgjvbroqtunleyzmdmfksabgw
icxjtbroqtunleyzkdmfksahww
icxjvbfoqtunleyzpamfqsahgw
icxjvbroknuyleyzpdmfksahgw
icxjvbroqtujleyzpdmaksaigw
icxjvbroqtnnlmyzpdmflsahgw
icxjvbroqtunlefzpdmfsfahgw
icxjvdroqtusleyzpdzfksahgw
icxjvbioqtunlsyzpdmfkshhgw
icxbvbrodtunleyzpdmoksahgw
icxjvbroqtuvleyzpdmfkbahmw
iyxjvbroqtunljvzpdmfksahgw
icxjvbroqtudleynddmfksahgw
icxjvwroqtnnleyzpdmfksahgz
ichjvbroqtunleyzjdmeksahgw
icxjvbrostunluyrpdmfksahgw
icfjvbroqtunleyxpdgfksahgw
nhxjvbroqtunlerzpdmfksahgw
icxjvbrothunlexzpdmfksahgw
icxjvbrzltqnleyzpdmfksahgw
icxjvbrhqtunleyzpdmfksajgy
vcxjvjroqiunleyzpdmfksahgw
icxjfbroltunleyzpdmqksahgw
icxbvbroqtunleyzpdofasahgw
icxjvbkoqtunveyzpdmfksaqgw
icxsebroqtunleyzpdmuksahgw
icxjvbroquunlpyrpdmfksahgw
icxhvbroqtunjeyzpdmrksahgw
icdjvbroqtunlzyzpdmfksangw
jcxqvbroqtvnleyzpdmfksahgw
icxjvxroqtunleyrpdmfxsahgw
icxjvnroqtunleyzpdmfssyhgw
icxjvbraptunleyzpdofksahgw
icxjvbroatunleyjpdmfbsahgw
icxjvbroytlnlryzpdmfksahgw
iaxjvbroqkunleyzpdmfcsahgw
ucxjvbroqtuileyzzdmfksahgw
icxjqbroqtcnleyzpgmfksahgw
icxjvbloqtunleyzadmfksaqgw
icxjvbroqtunleyzkdmnksakgw
icxjvbroqtunleyjpdxfksahvw
iqxjvbroqtujleyzpdmfklahgw
icgjvbroqtunleyzpdmfksbhgb
icxjzbroqtunleyzpdmfkdahgg
icxjvbrobtunloywpdmfksahgw
icxavbroqtunleyfpdmfksahgd
icxjvbroqtunleyophmfksahkw
icxjndroqtunlyyzpdmfksahgw
icxjvbroqtjnleyppdmvksahgw
icxjvbroonfnleyzpdmfksahgw
icxjvbrqqtlnljyzpdmfksahgw
icxjvbrzqtunlelspdmfksahgw
icxjvbooqtunleyztdmfkfahgw
icajvbroltunlnyzpdmfksahgw
icxjvbroqtunleyzidmdkschgw
icxjvbroktupleyzpdmfksahyw
icxjcbroyeunleyzpdmfksahgw
icxjvbroqtunlezkpdmfksahsw
icxjvbroqtunlejzpcmfksrhgw
icxjvvroqtunlsyzkdmfksahgw
icxjnbroqtunbeyzpdmfpsahgw
itxjbbroqtunleyzpemfksahgw
icxjvbroqtunlrkzpdmfksshgw
rcxjvbroqtujlbyzpdmfksahgw
icxjvmroqtugleazpdmfksahgw
icxjvbfoqtunleylpdmfkeahgw
icnjvoroktunleyzpdmfksahgw
icxjvbroqtunlvyzpdmfkfahgr
icxjvbroqtgnseyzpdmfxsahgw
scxjvbroqtunleycpdmfksdhgw
icxhvbxoqtunleuzpdmfksahgw
icxjvbruqtunleyzpnmfksfhgw
icdjvbroqtunleyzpdmfksahup
ihxjvbroqtunleovpdmfksahgw
icxjvbroqtunleyzxdmfksahzv
ocxjvbioqtunleyzpdmfzsahgw
idxjvbroqtunlyyzpdofksahgw
izdjvbroqtunleyzpdtfksahgw
icxjvbrnqtunleyzpdmfksbhgb
icxjvbrjqtunleyhpdmrksahgw
icxjvbroqtunleyzpdbflsahgg
icxjvbmfqtunleyzpdmfkaahgw
idxjvbroqtunlsyzpdffksahgw
bcxjvbroqtunleyzpkmfkswhgw
ivxjvbroqtdnleyzpdmbksahgw
icxpvbboqtunleyzpdmfksahtw
ibxjvbroqtunlehzpdmfkmahgw
icxjvbboqtunleyzpdmfkaahgv
icxjlaroqtuileyzpdmfksahgw
icxjvbroftunleyzpdmfqsahew
ichjvbroqtunleyzpdmiwsahgw
icxrvbvoqtunleyzpdmiksahgw
icxjvbroqtunldydpdmfksahgl
icogvbroqtunleyzpdmfnsahgw
icxjvbroqtunleszodmfkswhgw
icxjvbrontunleyztemfksahgw
icxjvbrovtunleyzpdvkksahgw
icxjvbroqqucteyzpdmfksahgw
icmovbroptunleyzpdmfksahgw
icxjvbqoftunleyzvdmfksahgw
icxjvbdoqtunleyzpdmfkadhgw
icxjvbroqtunlgnzpdmfksaqgw
icxjvbroqtunieygpdyfksahgw
acdjvbroqtunleyzpdmfkwahgw
icxjvbroqtunleyzpdmfkfahlj
icxjvbgoqtunleyepdmvksahgw
icxjvbpobbunleyzpdmfksahgw
icxjvbroqtunleurpdmfktahgw
ipxjvbzoqtunleyzpdmffsahgw
icxjtbroqtunllyzpdmuksahgw
icxjvbroqtunbsyzadmfksahgw
ihxjvoroqtuqleyzpdmfksahgw
idxjmbroqqunleyzpdmfksahgw
wcxjvbdoqtunleyzpdmfksahgr
icxjvbroqtunleygptmfksahgj
ipxjvbrsqtunleyzpdmfksghgw
ycxjvbroqtunluyzkdmfksahgw
icxjvbroxtuulejzpdmfksahgw
icqjvbroqtunlwyypdmfksahgw
ioxjhbroqtunleyzphmfksahgw
icxjvbgoqnunleyzpdmfksahaw
mcxjvbroqtunleyzpdmfksihgh
icxjsbroqtunlqyzpdmfksawgw
icxjvbroqtuoleycpdmftsahgw
icxjvbroqtunleyzgdifksahlw
icxjvbmoqtunleyzjfmfksahgw
icxjvbroqtunlezopdmfksahge
icxjvbroqtbnlefzpdmfosahgw
tcxjvbromtunlevzpdmfksahgw
irxjgbroqtunleyzpdmfksthgw
icxjvbrojtunleyxpdmoksahgw
icxrvbroytpnleyzpdmfksahgw
icxjvbroqtunfeyupdmfksasgw
ihqjvbroqtunleyzpdmftsahgw
icxjobroqkunleozpdmfksahgw
icjjjbroqtualeyzpdmfksahgw
icxjvbroqtunaeytpdmfksahvw
icxjvbroqtunzeyzpdmfkshhxw
icxqvbroqtucleyzxdmfksahgw
icxjvbrogturleyzxdmfksahgw
icxjvoqoqtunleyzpdcfksahgw
iuxjvbroqtunleyzpdmfksopgw
icxjveroqtunleyzptmfksalgw
icxjvbroqtunleyzpdmfwcahhw
iwxjvbroqtlnleyzpdmfksyhgw
ectjvbroqtanleyzpdmfksahgw
icxjvnroqtyhleyzpdmfksahgw
icvjvhboqtunleyzpdmfksahgw
icxjtbroqtuzleyupdmfksahgw
icjjvproqtunleyzpsmfksahgw
icdjvbroqtutleyzpdmiksahgw
icxjvwroqtujleyzpdmfksahgc
icxjxbroqtunleyzpdwhksahgw
icxjvbqoqtunleyzpdmvfsahgw
icajvbroqtusleyzpdmfksaagw
icxjvbroqtunbtyzpdmfksmhgw
kcxjvbroqtxnleyzpdmfkskhgw
icxjvbqogfunleyzpdmfksahgw
icxjvbroqtubleyzpdmfdswhgw
icxjvprlqtunleyzpdmffsahgw
icxjxbroqtucleyzpdmfksakgw
dcxrvbroqtunleycpdmfksahgw
icxjvbrobtunleyzpomfksahgu
ocxrvbroqtunleyzpdmfssahgw
icxjvbroktunlejzpdmfksahzw
icxjvbrovtunleyzmdmfkhahgw
icxjvbroqtudleygpdmfksfhgw
bcxjvbroqtubllyzpdmfksahgw
icxwvbrontunzeyzpdmfksahgw
icxjvbroqtunleysjbmfksahgw
icxjvvroztunleyzpdmfksjhgw
ivxjxbroqtunleyzpdmfksahew
icxjvbroqtunleyupqufksahgw
icxjvmrcqtunleyzpdmxksahgw
icxjvgroqtunleyzpdgfkuahgw
icxjvbroqthnqeyfpdmfksahgw
icxjsbuodtunleyzpdmfksahgw
iuxjzbroqtunleyzpdrfksahgw
icxjvbrobtunlelzpdmfksahgs
icxjvbroqtzhljyzpdmfksahgw
inxtvbroqtunleyzpdmeksahgw
icgjvbroqtunleyztdmfksahgq
icxjvagoqtugleyzpdmfksahgw
icxuvbroqtunleyzpimfkyahgw
icxzvbroqtfhleyzpdmfksahgw
icxjjbroqtqnleyzpdmnksahgw
icjrvbroqtunleszpdmfksahgw
iexjvbroqtunlgyzpdmfksacgw
rcxjvbkoqtuoleyzpdmfksahgw
icxjvbroqgunlwyzpdmfksqhgw
icxjvbroqtunleqzpsmfksqhgw
icxjvbroqtubaeyzpdmfksaugw`}
