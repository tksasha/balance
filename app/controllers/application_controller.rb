# frozen_string_literal: true

class ApplicationController < ActionController::Base
  include ActsAsController

  protect_from_forgery

  # rubocop:disable Rails/LexicallyScopedActionFilter:
  before_action :initialize_resource, only: :new

  before_action :build_resource, only: :create
  # rubocop:enable Rails/LexicallyScopedActionFilter:

  helper_method :collection, :resource, :serializer

  # TODO: spec me
  def default_url_options
    { currency: Currency.parse(params[:currency]) }
  end

  def dashboard
    ::Frontend::Dashboard.new(params)
  end

  helper_method :dashboard
end
