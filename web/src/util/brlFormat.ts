export function formatBRL(value: number) {
  return Intl.NumberFormat('pt-BR', {
    style: 'currency',
    currency: 'BRL',
  }).format(value || 0);
}

export function centsToDecimal(cents: number | undefined) {
  if (!cents) return 0;
  return Math.round(cents / 100);
}

export function decimalToCents(decimal: number) {
  if (!decimal) return 0;
  return Math.round(decimal * 100);
}

export function centsToCurrency(...cents: (number | undefined)[]) {
  const decimal = centsToDecimal(
    cents.reduce((a, b) => {
      return (a ?? 0) + (b ?? 0);
    }, 0)
  );
  return decimal.toLocaleString('pt-BR', {
    style: 'currency',
    currency: 'BRL',
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  });
}
