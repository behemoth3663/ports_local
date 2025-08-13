# $FreeBSD: head/sysutils/terraform/Makefile 502401 2019-05-23 23:06:04Z lifanov $

PORTNAME=	terraform
PORTVERSION=	0.11.15
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local sysutils
PKGNAMESUFFIX?=	-${PORTVERSION:R}

MAINTAINER=	none@freebsd.org
COMMENT=	Provides a common configuration to launch infrastructure

LICENSE=	MPL20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

PLIST_FILES=	bin/${PORTNAME}-${PORTVERSION}
PLIST_FILES+=	bin/${PORTNAME}${PKGNAMESUFFIX}

JQ_GETVER_EXPR=	[.[].name | gsub("^[a-z]+";"") | select(test("^${PORTVERSION:C/[^.]+$/[0-9]+/:S/./\\\\./g}$$"; "n"))][0]

post-patch: .PHONY
	${REINPLACE_CMD} 's|Prerelease = "dev"|Prerelease = ""|' \
		${WRKSRC}/version/version.go

post-install: .PHONY
	${MV} ${STAGEDIR}${PREFIX}/bin/${PORTNAME} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION}
	${INSTALL} -l rs ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}${PKGNAMESUFFIX}

.include <bsd.port.mk>
