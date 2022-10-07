#

PORTNAME=	ecctl
PORTVERSION=	1.8.0
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	The official ESS and ECE command line interface

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	elastic

PLIST_FILES=	bin/${PORTNAME}

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

#	version string
#	commit  string
#	built   string
post-patch: .PHONY
	${REINPLACE_CMD} -r \
		-e 's|^([[:space:]]*version[[:space:]]+string).*$$|\1 = "${DISTVERSIONPREFIX}${PORTVERSION}"|' \
		-e 's|^([[:space:]]*commit[[:space:]]+string).*$$|\1 = "${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}"|' \
		-e 's|^([[:space:]]*built[[:space:]]+string).*$$|\1 = "${%FT%TZ:L:gmtime}"|' \
	${WRKSRC}/main.go

.include <bsd.port.mk>
