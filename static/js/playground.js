/**
 * Playground.js
 */


function init(argument) {
  // editor init
  // $('#code').linedtextarea();

  var editor = ace.edit("wrap");
    editor.setTheme("ace/theme/monokai");
    editor.getSession().setMode("ace/mode/golang");

  var Range = ace.require('ace/range').Range;
  // Event Binding
  $('.run').click(function(e) {

    $('.output').text('Loading...')

    $('.ace_gutter-cell').remove('ace_error');

    $.ajax({
      url: '/run',
      type: 'POST',
      data: {
        code:editor.getValue()
      },
      success: function(output) {
        if (output.error) {
          $('.output').text(output.error);

          var lins = output.error.split('\n')
          for(var i = 0;i < lins.length; i++){
           var errorline = lins[i].match(/\:[0-9]*\:/)
           if(errorline){
             var num = errorline[0].replace(/\:/g,"")
             $(".ace_gutter-cell:contains('"+num+"')").addClass('ace_error');
           }
          }
        } else {
          $('.output').text(output.message);
        }
      }
    })
  });


}

init();
