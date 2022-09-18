# frozen_string_literal: true

class ApplicationController < ActionController::Base
  include ActsAsController

  protect_from_forgery

  helper_method :collection, :resource, :serializer

  # TODO: spec me
  def default_url_options
    { currency: Currency(params[:currency]) }
  end
end
