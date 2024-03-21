/*
%%%%%%%%%%%%%插件版本：SG 1.0.0%%%%%%%%%%%%%
%%%%%%%%%%%%%该文件由工具自动生成，请不要手动修改%%%%%%%%%%%%%
%%%%%%%%%%%%%生成时间:2024-03-21 16:47:31%%%%%%%%%%%%%
*/

package gamedata

type Ball struct {
	Id	int	"index"	//唯一id
	Functionname	string	//道具名称
	Types	int	//对应方块类型
	Probability	int	//触发概率
	Addweighta	int	//提升道具a掉落概率
	Addweightb	int	//提升道具b掉落概率
	Addweightc	int	//提升道具c掉落概率
}
