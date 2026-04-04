{
  buildGoModule,
  useCnMirror ? false,
  lib,
}:
buildGoModule (finalAttrs: {
  pname = "bili-danmaku-tui";
  version = "0.1.3";

  src = ./.;

  vendorHash = "sha256-CdGMkJz2oxrxqiCYtUvms9J9ISekzVo+Aj6WSEPqYuk=";

  preBuild =
    if useCnMirror
    then ''
      export GOPROXY=https://goproxy.cn,direct
    ''
    else "";

  ldflags = [
    "-s"
  ];

  meta = {
    description = "终端展示bilibili弹幕";
    homepage = "https://github.com/Youthdreamer/bili-danmaku-tui";
    license = lib.licenses.mit;
    mainProgram = "bili-danmaku-tui";
  };
})
