function fun(){
  if ($("#follow-button").text() == "+ Follow"){
      // *** State Change: To Following ***
      // We want the button to squish (or shrink) by 10px as a reaction to the click and for it to last 100ms
      // $("#follow-button").animate({ width: '-=10px' }, 100, 'easeInCubic', function () {});

      // then now we want the button to expand out to it's full state
      // The left translation is to keep the button centred with it's longer width
      $("#follow-button").css("background-color", "green");
        $("#follow-button").text("Following");
    }else{

      // *** State Change: Unfollow ***
      // Change the button back to it's original state
      $("#follow-button").css("background-color", "#3399FF");
        $("#follow-button").text("+ Follow");
        // $("#follow-button").css("color", "#3399FF");
        // $("#follow-button").css("border-color", "#3399FF");
    }
}
$(function() {


    /*--------------------------------------
     #User Object
     --------------------------------------*/

    var User = {
        handle : '@apteharsh',
        img : 'user-512.webp',
    };

    /*--------------------------------------
     #State Management
     --------------------------------------*/

    $( 'main' ).on( 'click', 'textarea', function() {
        $( this ).parents( 'form' ).addClass( 'expand' );
    } );

    $( '.tweets' ).on( 'click', '.thread > .tweet', function() {
        $( this ).parents( '.thread' ).toggleClass( 'expand' );
    } );

    /*--------------------------------------
     #Templating
     --------------------------------------*/

    /**
     * Compile Templates
     */
    var tweet   = Handlebars.compile( $( '#template-tweet' ).html() );
    var compose = Handlebars.compile( $( '#template-compose' ).html() );
    var thread  = Handlebars.compile( $( '#template-thread' ).html() );

    /**
     * Create New Tweet
     */
    function renderTweet( User, message ) {
        var data = {
            handle : User.handle,
            img : User.img,
            message : message
        };
        return tweet( data );
    };

    /**
     * Compose Area
     */
    function renderCompose() {
        return compose();
    }

    /**
     * Create a New Thread
     */
    function renderThread( User, message ) {
        var getTweet   = renderTweet( User, message );
        var getCompose = renderCompose();

        var getThread = {
            tweetTemplate : getTweet,
            composeTemplate : getCompose
        };
        return thread( getThread );
    }

    /*--------------------------------------
     #Composition
     --------------------------------------*/

    $( document ).on( 'submit', 'form', function() {
        event.preventDefault();
        message = $( 'textarea', this ).val();

        if ( $( this ).parent( 'header').length ) {
            $( '.tweets' ).append( renderThread( User, message ) );
        } else {
            $( this ).parent().append( renderTweet( User, message ) );
        }

        $( 'textarea' ).val( '' );
        $( 'form' ).removeClass( 'expand' );
    } );

} );
