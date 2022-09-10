# Created by: Alan Eldridge <alane@FreeBSD.org>

PORTNAME=	firefox
DISTVERSION=	102.2.0
#PORTREVISION=	0
#PORTEPOCH=	2
CATEGORIES=	www
MASTER_SITES=	MOZILLA/${PORTNAME}/releases/${DISTVERSION}esr/source \
		MOZILLA/${PORTNAME}/candidates/${DISTVERSION}esr-candidates/build1/source
PKGNAMESUFFIX=	-esr
DISTFILES=	${DISTNAME}esr.source${EXTRACT_SUFX}

MAINTAINER=	gecko@FreeBSD.org
COMMENT=	Web browser based on the browser portion of Mozilla

# Attention: 2022Q2 is still on icu 70 and uses the embedded ICU
BUILD_DEPENDS=	nspr>=4.32:devel/nspr \
		nss>=3.79:security/nss \
		icu>=71.1:devel/icu \
		libevent>=2.1.8:devel/libevent \
		harfbuzz>=4.1.0:print/harfbuzz \
		graphite2>=1.3.14:graphics/graphite2 \
		png>=1.6.37:graphics/png \
		dav1d>=1.0.0:multimedia/dav1d \
		libvpx>=1.11.0:multimedia/libvpx \
		${PYTHON_PKGNAMEPREFIX}sqlite3>0:databases/py-sqlite3@${PY_FLAVOR} \
		v4l_compat>0:multimedia/v4l_compat \
		autoconf2.13:devel/autoconf2.13 \
		nasm:devel/nasm \
		yasm:devel/yasm \
		zip:archivers/zip

USE_GECKO=	gecko
CPE_PRODUCT=	${PORTNAME}
CONFLICTS_INSTALL=	firefox-esr
USE_MOZILLA=	-sqlite
CFLAGS_powerpc64le=	-DSQLITE_BYTEORDER=1234
# work around bindgen not finding ICU, e.g.
# dist/include/mozilla/intl/ICU4CGlue.h:8:10: fatal error: 'unicode/uenum.h' file not found, err: true
CONFIGURE_ENV+=	BINDGEN_CFLAGS="-I${LOCALBASE}/include"

USES=		tar:xz

FIREFOX_ICON=		${MOZILLA}.png
FIREFOX_ICON_SRC=	${PREFIX}/lib/${MOZILLA}/browser/chrome/icons/default/default48.png
FIREFOX_DESKTOP=	${MOZSRC}/taskcluster/docker/${MOZILLA}-snap/${MOZILLA}.desktop
MOZ_OPTIONS=	--enable-application=browser \
		--enable-official-branding

.include "${.CURDIR}/../../www/firefox/Makefile.options"

.include <bsd.port.options.mk>

.if ${ARCH} == powerpc64
MOZ_OPTIONS+=	--disable-webrtc --without-wasm-sandboxed-libraries
.else
BUILD_DEPENDS+=	${LOCALBASE}/share/wasi-sysroot/lib/wasm32-wasi/libc++abi.a:devel/wasi-libcxx \
		${LOCALBASE}/share/wasi-sysroot/lib/wasm32-wasi/libc.a:devel/wasi-libc \
		${LOCALBASE}/llvm${LLVM_DEFAULT}/lib/clang/${LLVM_VERSION}/lib/wasi/libclang_rt.builtins-wasm32.a:devel/wasi-compiler-rt${LLVM_DEFAULT}
MOZ_OPTIONS+=	--with-wasi-sysroot=${LOCALBASE}/share/wasi-sysroot
.endif

post-patch:
	@${REINPLACE_CMD} -e 's/%u/%U/' -e '/X-MultipleArgs/d' \
		-e '/^Icon/s/=.*/=${FIREFOX_ICON:R}/' \
		${FIREFOX_DESKTOP}
	@${REINPLACE_CMD} -e 's|%%LOCALBASE%%|${LOCALBASE}|g' \
		${WRKSRC}/browser/app/nsBrowserApp.cpp

pre-configure:
	(cd ${WRKSRC} && ${LOCALBASE}/bin/autoconf2.13)
	(cd ${WRKSRC}/js/src/ && ${LOCALBASE}/bin/autoconf2.13)

post-install:
	${INSTALL_DATA} ${FIREFOX_DESKTOP} ${STAGEDIR}${PREFIX}/share/applications/
	${MKDIR} ${STAGEDIR}${PREFIX}/share/pixmaps
	${LN} -sf ${FIREFOX_ICON_SRC} ${STAGEDIR}${PREFIX}/share/pixmaps/${FIREFOX_ICON}

.include <bsd.port.pre.mk>

LLVM_DEFAULT=13
LLVM_VERSION=13.0.1
DEFAULT_VERSIONS:=${DEFAULT_VERSIONS:C/llvm=.*/llvm=13/}

OPT_LEVEL=3

.if ! ${CFLAGS:M-fomit-frame-pointer}
  CFLAGS+=-fomit-frame-pointer
.endif

CFLAGS:=${CFLAGS:N-O*} -O${OPT_LEVEL}

.if !empty(CPUTYPE)
  CFLAGS+=-march=${CPUTYPE}
.endif

#MOZ_OPTIONS+=	--disable-maintenance-service
MOZ_OPTIONS+=	--disable-crashreporter
MOZ_OPTIONS+=	--disable-parental-controls
MOZ_OPTIONS+=	--disable-necko-wifi
MOZ_OPTIONS+=	--disable-rust-simd
#?MOZ_OPTIONS:=	${MOZ_OPTIONS:S/--enable-lto=thin/--enable-lto=full/}

MOZ_EXPORT:=	${MOZ_EXPORT:NLLVM_CONFIG=llvm-config11}
MOZ_EXPORT:=	${MOZ_EXPORT:C,^(MAKE|PKG_CONFIG)=,\1=/usr/local/bin/,}
MOZ_EXPORT:=	${MOZ_EXPORT:NRUSTFLAGS=*} RUSTFLAGS="-C opt-level=${CFLAGS:M-O*:S/-O//:S/3/2/} -C target-cpu=${CFLAGS:M-march=*:S/-march=//:S/^$/x86-64/}"

MOZ_EXPORT+=	MOZ_TELEMETRY_ON_BY_DEFAULT=0
MOZ_EXPORT+=	MOZ_DATA_REPORTING=0
MOZ_EXPORT+=	MOZ_SCTP=0

.include <bsd.port.post.mk>
