
func plusMinus(arr []int32) {
    var (
        p float32=0
        n float32=0
        z float32=0
        l = len(arr)
    )
    
    for i:=0; i<l;i++{
        if arr[i] > 0{
            p++
        } else if arr[i] < 0{
            n++
        } else{
            z++
        }
    }
    fmt.Println(p/float32(l))
    fmt.Println(n/float32(l))
    fmt.Println(z/float32(l))
}
