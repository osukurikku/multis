const modsString = [
  'NF',
  'EZ',
  'NV',
  'HD',
  'HR',
  'SD',
  'DT',
  'RX',
  'HT',
  'NC',
  'FL',
  'AU', // Auto.
  'SO',
  'AP', // Autopilot.
  'PF',
  'K4',
  'K5',
  'K6',
  'K7',
  'K8',
  'K9',
  'RN', // Random
  'LM', // LastMod. Cinema?
  'K9',
  'K0',
  'K1',
  'K3',
  'K2'
]

export default m => {
  var r = []

  // has nc => remove dt
  if ((m & 512) === 512) {
    m = m & ~64
  }
  // has pf => remove sd
  if ((m & 16384) === 16384) {
    m = m & ~32
  }

  modsString.forEach(function (v, idx) {
    var val = 1 << idx
    if ((m & val) > 0) {
      r.push(v)
    }
  })
  if (r.length > 0) {
    return '+ ' + r.join(', ')
  }
  return ''
}
