# frozen_string_literal: true

RSpec.shared_examples 'destroy.js' do
  it { is_expected.to render_template :destroy }

  it { expect(response.content_type).to eq 'text/javascript; charset=utf-8' }
end
