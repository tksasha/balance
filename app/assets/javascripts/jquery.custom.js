jQuery.fn.custom = function(f) {
  if(typeof f === 'function') {
    f.call();
  }

  return(this);
}
