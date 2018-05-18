func (jp JProxy) Less(i, j int) bool {
    return jp[i].GetName()+" - "+jp[i].GetLocal() < jp[j].GetName()+" - "+jp[j].GetLocal()
}
