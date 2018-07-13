package old

func main()  {
	var i I
	var t *T
	i = t
	describe(i)
	i.M()

	i=&T{"hello"}
	describe(i)
	i.M()
}