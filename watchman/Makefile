PORTNAME=	watchman
PORTVERSION=	4.9.0
DISTVERSIONPREFIX=	v
PORTREVISION=	3
PORTEPOCH=	1
CATEGORIES=	local sysutils
PKGNAMESUFFIX=	-local

MAINTAINER=	yuri@FreeBSD.org
COMMENT=	File alteration monitoring service
WWW=		https://facebook.github.io/watchman/

LICENSE=	APACHE20

BROKEN_mips=	fails to build: watchman.h:169: undefined reference to '__sync_add_and_fetch_8'
BROKEN_mips64=	fails to build: watchman.h:169: undefined reference to '__sync_add_and_fetch_8'

BUILD_DEPENDS=	${LOCALBASE}/include/sys/inotify.h:devel/libinotify

USES=		autoreconf gmake pkgconfig llvm:22,build
GNU_CONFIGURE=	yes
CONFIGURE_ENV=	ac_cv_prog_HAVE_ARC=false ac_cv_header_sys_inotify_h=no
CONFIGURE_ARGS=	--without-python --without-ruby --disable-cppclient # in separate ports
TEST_TARGET=	check

USE_GITHUB=	yes
GH_ACCOUNT=	facebook

PLIST_FILES=	bin/${PORTNAME} \
		"@dir(,,2777) /var/run/${PORTNAME}"
PORTDOCS=	README.markdown

OPTIONS_DEFINE=	BACKTRACE PCRE DOCS
OPTIONS_DEFAULT=BACKTRACE PCRE

BACKTRACE_DESC=		Stack backtrace support via (lib)execinfo
BACKTRACE_LIBS=		-lexecinfo

PCRE_LIB_DEPENDS=	libpcre.so:devel/pcre
PCRE_CONFIGURE_WITH=	pcre

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}^{}

CPPFLAGS+=	-Wno-error

CFLAGS:=	-Oz ${CFLAGS:N-O*}
CFLAGS+=	-DWATCHMAN_BUILD_INFO='\"${GIT_COMMIT:sh:M*:[1]}\"'
CFLAGS+=	-fno-unwind-tables -fno-asynchronous-unwind-tables

#-CFLAGS+=	-fomit-frame-pointer # -fno-omit-frame-pointer
#-LDFLAGS+=	-Wl,-O3

CFLAGS+=	-flto=full
LDFLAGS+=	-Wl,--lto-O3
CFLAGS+=	-fwhole-program-vtables

#CFLAGS+=	-fno-unique-section-names

CFLAGS+=	-faddrsig
LDFLAGS+=	-Wl,--icf=all

CFLAGS+=	-ffunction-sections
CFLAGS+=	-fdata-sections
LDFLAGS+=	-Wl,--gc-sections

LDFLAGS+=	-Wl,--as-needed
LDFLAGS+=	-Wl,--build-id=none
LDFLAGS+=	-Wl,--sort-common

PORTSCOUT=	ignore:1 # newer versions (YYYY.MM.DD.dd) use F_GETPATH that isn't currently available on FreeBSD

JQ_GETVER_EXPR=	[.[].name | gsub("^[a-z]+";"") | \
		select(test("^${PORTVERSION:C/[^.]+$/[0-9]+/:S/./\\\\./g}$$"; "n"))][0]

post-patch: .PHONY
# https://www.gnu.org/prep/standards/html_node/Directory-Variables.html
	@${REINPLACE_CMD} -e 's,xprefix/var,localstatedir,' \
		-e '\,/etc, { s//$$sysconfdir/; /=/s/^/eval /; }' \
		${WRKSRC}/configure.ac
	@${REINPLACE_CMD} -e '/^docdir = /d' \
		-e '/not-empty/d; /chmod g+s/d;' \
		${WRKSRC}/Makefile.am

.include <bsd.port.mk>
