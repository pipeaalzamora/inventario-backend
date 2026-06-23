#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

echo "== Sensitive service/facade methods =="
rg -n "func .*\\((ctx context\\.Context|.*context\\.Context).*\\).*(companyID|companyId|storeID|storeId|warehouseID|warehouseId|purchaseID|delivery.*ID|storeProductID|storeProductId)" domain/services domain/facades || true

echo
echo "== Ownership checks =="
rg -n "EveryPower\\(|SomePower\\(|PowerPrefixCompany|PowerPrefixStore|ExistsSupplierInCompany|store\\.CompanyID|company_id" domain/services domain/facades || true

echo
echo "== Direct repo calls from facades using tenant IDs =="
rg -n "\\.(Get|Create|Update|Delete|List|Assign|Unassign).*\\((ctx, )?.*(companyID|companyId|storeID|storeId|warehouseID|warehouseId|supplierID|purchaseID)" domain/facades || true
