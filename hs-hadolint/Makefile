PORTNAME=	hadolint
DISTVERSION=	2.15.1
#PORTREVISION=	1
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

#https://docs.freebsd.org/en/books/porters-handbook/special/index.html#using-cabal
# The file including the USE_CABAL. To regenerate it it's required to run:
# 1) make clean cabal-extract
# 2) make cabal-configure
# 3) make make-use-cabal | tee Makefile.cabal
# 4) make makesum

.include <bsd.port.mk>
