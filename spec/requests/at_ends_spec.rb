require 'rails_helper'

RSpec.describe 'AtEnds', type: :request do
  it_behaves_like 'show', '/at_end.js'
end
