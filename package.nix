{
  buildGoModule,
  useCnMirror ? false,
  lib,
}:
buildGoModule (finalAttrs: {
  pname = "bili-danmaku-tui";
  version = "0.1.2";

  src = ./.;

  vendorHash = "sha256-F6CxCj9bl48UMGsVekxp0lsIZT7y/Cfpc1wszPefuxY=";

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
    license = lib.licenses.MIT;
    mainProgram = "bili-danmaku-tui";
  };
})
