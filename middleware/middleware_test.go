package middleware

//
//var b = tele.New(tele.BotSettings{
//	OfflineMode: true,
//})
//
//func TestRecover(t *testing.T) {
//	onError := func(err error) {
//		require.Error(t, err, "recover test")
//	}
//
//	h := func(c tele.Context) error {
//		panic("recover test")
//	}
//
//	assert.Panics(t, func() {
//		h(nil)
//	})
//
//	assert.NotPanics(t, func() {
//		Recover(onError)(h)(nil)
//	})
//}
