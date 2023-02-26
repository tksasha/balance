# frozen_string_literal: true

class ApplicationController < ActionController::Base
  protect_from_forgery

  # TODO: spec me
  def default_url_options
    { currency: Currency.parse(params[:currency]) }
  end

  # TODO: spec me
  def dashboard
    ::Frontend::Dashboard.new(params)
  end

  helper_method :dashboard
end
