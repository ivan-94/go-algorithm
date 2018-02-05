package hash

// TODO:
// 这个文件实现了自动增长的Hash表
// go 原生的map类型就是自增长的：
// > go map类型每个bucket包含8个键值对。
// > 哈希值的低端位用于确定选择哪个bucket。
// > 每个bucket包含了一些hash值得高端位，用于区分每个数据项在
// > bucket中的存储;
// >
// > 如果bucket超过8个键值对, go会创建一个新的列表连接到这个bucket
// >
// > 当hash表增长时， go会分配一个新的buckets数组， 新的buckets是旧的buckets的
// > 两倍。 并使用增量的形式将旧的buckets数据拷贝到新的buckets中
