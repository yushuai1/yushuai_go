package 工厂模式

import "testing"

func Ttest(t *testing.T ) {

	var c WorkerCreator

	c = new(ProgrammerCreator)
	p := c.Create()
	taskProject := "Project"
	p.Work(&taskProject)

	c = new(FarmerCreator)
	f := c.Create()
	taskFarmland := "farmland"
	f.Work(&taskFarmland)
}
