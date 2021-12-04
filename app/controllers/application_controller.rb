# frozen_string_literal: true

class ApplicationController < ActionController::Base
  include ActsAsController

  protect_from_forgery

  helper_method :collection, :resource

  # TODO: spec me
  def default_url_options
    { currency: ParseCurrencyService.call(params[:currency]) }
  end
end
