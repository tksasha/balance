# frozen_string_literal: true

class BaseController < ApplicationController
  include ActsAsController

  # rubocop:disable Rails/LexicallyScopedActionFilter:
  before_action :initialize_resource, only: :new

  before_action :build_resource, only: :create
  # rubocop:enable Rails/LexicallyScopedActionFilter:

  helper_method :collection, :resource, :serializer
end
