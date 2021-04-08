# $FreeBSD: head/sysutils/terraform/Makefile 520926 2019-12-26 15:43:22Z lifanov $

PORTNAME=	terraform
PORTVERSION=	0.14.10
DISTVERSIONPREFIX=	v
CATEGORIES=	local
PKGNAMESUFFIX?=	-${PORTVERSION:R}

MAINTAINER=	none@freebsd.org
COMMENT=	Provides a common configuration to launch infrastructure

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

PLIST_FILES=	bin/${PORTNAME}-${PORTVERSION}
PLIST_FILES+=	bin/${PORTNAME}${PKGNAMESUFFIX}

JQ_GETVER_EXPR=	[.[].name | gsub("^[a-z]+";"") | select(test("^${PORTVERSION:C/[^.]+$/[0-9]+/:S/./\\\\./g}$$"; "n"))][0]

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor &&  make gomod-vendor >Makefile.deps && make makesum && mkcal
.-include "${.CURDIR}/Makefile.deps"

post-patch:
	${REINPLACE_CMD} 's|Prerelease = "dev"|Prerelease = ""|' \
		${WRKSRC}/version/version.go

post-install:
	${MV} ${STAGEDIR}${PREFIX}/bin/${PORTNAME} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION}
	${INSTALL} -l rs ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}${PKGNAMESUFFIX}

.include <bsd.port.mk>
