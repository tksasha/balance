# frozen_string_literal: true

RSpec.shared_examples 'show.js' do
  it { is_expected.to render_template :show }

  it { expect(response.content_type).to eq 'text/javascript; charset=utf-8' }
end

RSpec.shared_examples 'show.html' do
  it { is_expected.to render_template :show }
end
