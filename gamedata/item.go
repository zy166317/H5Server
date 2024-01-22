/*
%%%%%%%%%%%%%插件版本：SG 1.0.0%%%%%%%%%%%%%
%%%%%%%%%%%%%该文件由工具自动生成，请不要手动修改%%%%%%%%%%%%%
%%%%%%%%%%%%%生成时间:2023-12-14 09:51:41%%%%%%%%%%%%%
*/

package gamedata

type Item struct {
	Id	int	"index"	//唯一id
	Functionname	string	//道具名称
	Icon	string	//图标
	Price	[]int	//消耗道具
	Args1	[][]int	//参数一
	Args2	float32	//参数二
}
