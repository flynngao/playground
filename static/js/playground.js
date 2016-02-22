/**
 * Playground.js
 */


function init(argument) {
  // editor init
  $('#code').linedtextarea();


  // Event Binding
  $('.run').click(function(e) {
    $.ajax({
      url: '/run',
      type: 'POST',
      data: {
        code:$('#code').val()
      },
      success: function(output) {
        if (output.error) {
          $('.output').text(output.error);
        } else {
          $('.output').text(output.message);
        }
      }
    })
  });


}

init();
