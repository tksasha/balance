# frozen_string_literal: true

module Frontend
  class ApplicationController < ApplicationController
    def default_url_options
      { currency: Currency.parse(params[:currency]) }
    end
  end
end
