# frozen_string_literal: true

module Permitter
  def acp(params)
    ActionController::Parameters.new params
  end
end
