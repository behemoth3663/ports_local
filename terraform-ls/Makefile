#

PORTNAME=	terraform-ls
PORTVERSION=	0.27.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Terraform Language Server

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

#GO_BUILDFLAGS+=	-ldflags="-s -w -X 'github.com/main.version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

#rm -f -- distinfo Makefile.deps && doas make clean && make makesum gomod-vendor && make gomod-vendor >Makefile.deps && doas make clean && make makesum package && instpc
.-include "${.CURDIR}/Makefile.deps"

#var version = "0.0.0"
#var prerelease = "dev"
post-patch:
	${REINPLACE_CMD} -r \
		-e 's|^([[:space:]]*var[[:space:]]+version[[:space:]]*=[[:space:]]*)".*"[[:space:]]*$$|\1"${PORTVERSION}"|' \
		-e 's|^([[:space:]]*var[[:space:]]+prerelease[[:space:]]*=[[:space:]]*)".*"[[:space:]]*$$|\1""|' \
	${WRKSRC}/version.go

.include <bsd.port.mk>
