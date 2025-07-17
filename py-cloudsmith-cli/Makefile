PORTNAME=	cloudsmith-cli
PORTVERSION=    1.8.3
CATEGORIES=	local devel python
MASTER_SITES=	PYPI
PKGNAMEPREFIX=	${PYTHON_PKGNAMEPREFIX}
DISTNAME=	${PORTNAME:S/-/_/g}-${PORTVERSION}

MAINTAINER=	none@freebsd.org
COMMENT=	Cloudsmith Command Line Interface
WWW=		https://github.com/cloudsmith-io/cloudsmith-cli

LICENSE=	APACHE20

RUN_DEPENDS=	${PYTHON_PKGNAMEPREFIX}click>=7.0:devel/py-click@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}click-configfile>=0.2.3:local/py-click-configfile@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}click-didyoumean>=0.0.3:devel/py-click-didyoumean@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}click-spinner>=0.1.7:local/py-click-spinner@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}cloudsmith-api>=2.0.18:local/py-cloudsmith-api@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}keyring>=25.4.1:security/py-keyring@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}requests>=2.18.4:www/py-requests@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}requests-toolbelt>=0.8.0:www/py-requests-toolbelt@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}semver>=2.7.9:devel/py-semver@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}urllib3>0:net/py-urllib3@${PY_FLAVOR}

USES=		python
USE_PYTHON=	autoplist concurrent pep517

post-extract: .PHONY
	echo ${PORTVERSION} > ${WRKSRC}/VERSION

.include <bsd.port.mk>
