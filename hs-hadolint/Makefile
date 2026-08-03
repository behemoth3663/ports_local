PORTNAME=	hadolint
DISTVERSION=	2.15.1
PORTREVISION=	1
CATEGORIES=	devel haskell
PKGNAMESUFFIX=	-local

MAINTAINER=	tiga@FreeBSD.org
COMMENT=	Dockerfile linter, validate inline bash, written in Haskell
WWW=		https://github.com/hadolint/hadolint

LICENSE=	GPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cabal
BUILD_ARGS+=	--allow-newer=time
BUILD_ARGS+=	--enable-executable-static
BUILD_ARGS+=	--enable-split-sections
BUILD_ARGS+=	--disable-library-profiling
BUILD_ARGS+=	--disable-executable-profiling
BUILD_ARGS+=	--ghc-options="-O2"
BUILD_ARGS+=	--ghc-options="-optl=-s"
BUILD_ARGS+=	--ghc-options="-flate-dmd-anal"
#BUILD_ARGS+=	--ghc-options="-fllvm"
#BUILD_ARGS+=	--ghc-options="-funbox-strict-fields"
#BUILD_ARGS+=	--ghc-options="-rtsopts=none"

#https://docs.freebsd.org/en/books/porters-handbook/special/index.html#using-cabal
# The file including the USE_CABAL. To regenerate it it's required to run:
# 1) make clean cabal-extract
# 2) make cabal-configure
# 3) make make-use-cabal | tee Makefile.cabal
# 4) make makesum

.include <bsd.port.mk>
