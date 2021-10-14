#

PORTNAME=	terraform
PORTVERSION=	1.0.9
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils
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

GO_BUILDFLAGS=	-ldflags='-s -w -extldflags "-static"'

JQ_GETVER_EXPR=	[.[].name | gsub("^[a-z]+";"") | select(test("^${PORTVERSION:C/[^.]+$/[0-9]+/:S/./\\\\./g}$$"; "n"))][0]

post-install: .PHONY
	${MV} ${STAGEDIR}${PREFIX}/bin/${PORTNAME} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION}
	${INSTALL} -l rs ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}${PKGNAMESUFFIX}

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
##GO_BUILDFLAGS+=-ldflags="-s -w -X github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Prerelease="
#GO_ENV+=	GOOS=${OPSYS:tl}
#.include <bsd.port.pre.mk>
#GO_CMD=	${LINUXBASE}/usr/local/go/bin/go
#GO_BUILDFLAGS:=${GO_BUILDFLAGS:N-mod=vendor} -mod=mod
#.include <bsd.port.post.mk>
