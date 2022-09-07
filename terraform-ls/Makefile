#

PORTNAME=	terraform-ls
PORTVERSION=	0.29.2
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Terraform Language Server

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

PLIST_FILES=	bin/${PORTNAME}

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

#var version = "0.0.0"
#var versionPrerelease = "dev"
post-patch:
	${REINPLACE_CMD} -r \
		-e 's|^([[:space:]]*var[[:space:]]+version[[:space:]]*=[[:space:]]*)".*"[[:space:]]*$$|\1"${PORTVERSION}"|' \
		-e 's|^([[:space:]]*var[[:space:]]+versionPrerelease[[:space:]]*=[[:space:]]*)".*"[[:space:]]*$$|\1""|' \
	${WRKSRC}/version.go

.include <bsd.port.mk>
