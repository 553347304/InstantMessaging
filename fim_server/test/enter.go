
func (l *AuthenticationLogic) Authentication(req *types.AuthenticationRequest) (resp string, err error) {
	// todo: add your logic here and delete this line

	if algorithms.InList(l.svcCtx.Config.WhiteList, req.VaildPath) {
		logs.Info("白名单", req.VaildPath)
		return "ok", nil
	}

	claims := jwts.ParseToken(req.Token)
	if claims == nil {
		err = errors.New("认证失败: " + err.Error())
		return
	}
	_, err = l.svcCtx.Redis.Get(fmt.Sprintf("logout_%s", req.Token)).Result()
	if err != nil {
		logs.Error("在黑名单中")
		err = errors.New("认证失败: " + err.Error())
		return
	}

	return "ok", nil
}
