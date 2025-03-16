export function formatBRL(value: number) {
  return Intl.NumberFormat('pt-BR', {
    style: 'currency',
    currency: 'BRL',
  }).format(value);
}

export function centsToDecimal(cents: number) {
  if (!cents) return 0;
  return cents / 100;
}

export function decimalToCents(decimal: number) {
  return Math.round(decimal * 100);
}

export function centsToCurrency(...cents: number[]) {
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
