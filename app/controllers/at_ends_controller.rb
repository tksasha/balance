# frozen_string_literal: true

class AtEndsController < ApplicationController
  private

  def resource
    @resource ||= AtEndService.new params
  end
end
