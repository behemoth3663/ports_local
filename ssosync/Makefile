PORTNAME=	ssosync
PORTVERSION=	2.3.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Populate AWS SSO directly with Google Apps users
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

BUILD_DEPENDS=	git:devel/git

USES=		go:1.25,modules

USE_GITHUB=	yes
GH_ACCOUNT=	awslabs

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.commit=$$(${GIT_COMMIT})' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.date=${%FT%TZ:L:gmtime}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.builtBy=${OPSYS}'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|}

.include <bsd.port.mk>
