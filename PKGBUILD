# Maintainer: NXShock <nxshock@gmail.com>
pkgname=promodj
pkgver=0.0.1
pkgrel=0
pkgdesc="Proxy client for promodj.com"
arch=('x86_64' 'aarch64')
license=('MIT')
url='https://github.com/nxshock/$pkgname'
depends=('ffmpeg')
makedepends=('go' 'git')
options=("!strip")
backup=("etc/$pkgname.toml")
source=("git+https://github.com/nxshock/$pkgname.git")
sha256sums=('SKIP')

build() {
    cd "$srcdir/$pkgname"
	export CGO_CPPFLAGS="${CPPFLAGS}"
	export CGO_CFLAGS="${CFLAGS}"
	export CGO_CXXFLAGS="${CXXFLAGS}"
	export CGO_LDFLAGS="${LDFLAGS}"
    go build -o $pkgname -buildmode=pie -trimpath -ldflags="-linkmode=external -s -w"
}

package() {
    cd "$srcdir/$pkgname"
    install -Dm755 "$pkgname"          "$pkgdir/usr/bin/$pkgname"
    install -Dm644 "$pkgname.toml"     "$pkgdir/etc/$pkgname.toml"
    install -Dm644 "$pkgname.service"  "$pkgdir/usr/lib/systemd/system/$pkgname.service"
    install -Dm644 "$pkgname.sysusers" "$pkgdir/usr/lib/sysusers.d/$pkgname.conf"

	mkdir "$pkgdir/usr/lib/$pkgname"
    cp -R "site" "$pkgdir/usr/lib/$pkgname"
	cp -R "templates" "$pkgdir/usr/lib/$pkgname"
}
