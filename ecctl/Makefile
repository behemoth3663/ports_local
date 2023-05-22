PORTNAME=	ecctl
PORTVERSION=	1.9.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	The official ESS and ECE command line interface
WWW=		https://github.com/${GH_ACCOUNT}/${PORTNAME}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	elastic

PLIST_FILES=	bin/${PORTNAME}

GH_COMMIT=	git ls-remote --tags ${WWW} refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

#	version string
#	commit  string
#	built   string
post-patch: .PHONY
	${REINPLACE_CMD} -r \
		-e 's|^([[:space:]]*version[[:space:]]+string).*$$|\1 = "${DISTVERSIONPREFIX}${PORTVERSION}"|' \
		-e 's|^([[:space:]]*commit[[:space:]]+string).*$$|\1 = "${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}"|' \
		-e 's|^([[:space:]]*built[[:space:]]+string).*$$|\1 = "${%FT%TZ:L:gmtime}"|' \
	${WRKSRC}/main.go

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
