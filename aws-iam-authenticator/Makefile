PORTNAME=	aws-iam-authenticator
PORTVERSION=	0.7.7
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Use AWS IAM credentials to authenticate to a Kubernetes cluster
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes-sigs

GO_TARGET=	./cmd/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'sigs.k8s.io/${PORTNAME}/pkg.Version=${PORTVERSION}' \
	-X 'sigs.k8s.io/${PORTNAME}/pkg.BuildDate=${%FT%TZ:L:gmtime}' \
	-X 'sigs.k8s.io/${PORTNAME}/pkg.CommitID=${GIT_COMMIT:sh}'"

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@test ! -d ${WRKSRC}/vendor/github.com || ${LN} -s ${WRKSRC}/src/github.com/${GH_ACCOUNT} ${WRKSRC}/vendor/github.com
	@${MKDIR} ${WRKSRC}/vendor/go.yaml.in/yaml
	@test ! -e ${WRKSRC}/vendor/gopkg.in/yaml.v2 || ${RLN} ${WRKSRC}/vendor/gopkg.in/yaml.v2 ${WRKSRC}/vendor/go.yaml.in/yaml/v2
	@test ! -e ${WRKSRC}/vendor/gopkg.in/yaml.v3 || ${RLN} ${WRKSRC}/vendor/gopkg.in/yaml.v3 ${WRKSRC}/vendor/go.yaml.in/yaml/v3

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes GHT_PROJ=klog GHT_PATH=k8s.io/klog/v2 )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=json GHT_PATH=sigs.k8s.io/json )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=structured-merge-diff GHT_PATH=sigs.k8s.io/structured-merge-diff/v6 )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=randfill GHT_PATH=sigs.k8s.io/randfill )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

.sinclude "${.CURDIR}/generated.deps.mk"
.include <bsd.port.mk>
