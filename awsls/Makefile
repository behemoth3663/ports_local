#

PORTNAME=	awsls
PORTVERSION=	0.11.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	List command for AWS resources

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.md

RUN_DEPENDS=	git:devel/git

#go:modules
# Errors found during processing:
# error getting refs for aws/aws-sdk-go-v2: apis.get https://api.github.com/repos/aws/aws-sdk-go-v2/git/refs/tags/v1.1.1: 403, body: {
#  "documentation_url": "https://docs.github.com/en/free-pro-team@latest/rest/overview/resources-in-the-rest-api#secondary-rate-limits",
#  "message": "You have exceeded a secondary rate limit. Please wait a few minutes before you try again."
#}
USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	jckuester

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/internal.version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/internal.commit=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/internal.date=${%FT%TZ:L:gmtime}'"

#without GOPROXY=off
do-build: .PHONY
	(cd ${GO_WRKSRC} && \
		${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} build ${GO_BUILDFLAGS} \
			-o ${GO_WRKDIR_BIN}/${PORTNAME})

#.-include "${.CURDIR}/Makefile.deps"
#.include <bsd.port.mk>
.include <bsd.port.pre.mk>
GO_ENV:=	${GO_ENV:NGO111MODULE=*}
.include <bsd.port.post.mk>
